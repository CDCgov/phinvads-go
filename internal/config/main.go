package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Addr       *string
	Dsn        *string
	TlsEnabled *bool
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func LoadConfig() *Config {
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbString := fmt.Sprintf(`postgresql://%s@%s:%s/%s`, dbUser, dbHost, dbPort, dbName)

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	addrString := fmt.Sprintf(`%s:%s`, host, port)

	tlsEnabledString := os.Getenv("TLS_ENABLED")
	tlsEnabled, err := strconv.ParseBool(tlsEnabledString)
	if err != nil {
		tlsEnabled = true
	}

	addr := flag.String("addr", addrString, "HTTP network address")
	dsn := flag.String("dsn", dbString, "PostgreSQL data source name")
	tls := flag.Bool("tls", tlsEnabled, "Enable TLS")
	flag.Parse()

	return &Config{
		Addr:       addr,
		Dsn:        dsn,
		TlsEnabled: tls,
	}
}
