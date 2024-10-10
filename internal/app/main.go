package app

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	cfg "github.com/CDCgov/phinvads-go/internal/config"
	"github.com/CDCgov/phinvads-go/internal/database"
	rp "github.com/CDCgov/phinvads-go/internal/database/models/repository"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Application struct {
	logger     *slog.Logger
	db         *sql.DB
	server     *http.Server
	repository *rp.Repository
	tlsEnabled bool
}

func SetupApp(cfg *cfg.Config) *Application {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := database.CreateDB(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	rp := rp.NewRepository(db)

	app := &Application{
		logger:     logger,
		db:         db,
		repository: rp,
		tlsEnabled: *cfg.TlsEnabled,
	}

	srv := &http.Server{
		Addr:         *cfg.Addr,
		Handler:      app.routes(),
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	app.server = srv
	return app
}

// Run starts the application
func (app *Application) Run() {
	app.logger.Info("starting server", slog.String("addr", app.server.Addr))

	var err error
	if app.tlsEnabled {
		tlsCert := getEnvWithFallback("TLS_CERT", "./tls/localhost.pem")
		tlsKey := getEnvWithFallback("TLS_KEY", "./tls/localhost-key.pem")

		err = app.server.ListenAndServeTLS(tlsCert, tlsKey)
	} else {
		err = app.server.ListenAndServe()
	}

	app.logger.Error(err.Error())

	defer app.db.Close()

	os.Exit(1)
}

// getEnvWithFallback retrieves the value of an environment variable by key;
// if no value is found, the provided fallback value is returned
func getEnvWithFallback(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
