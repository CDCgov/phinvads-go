package main

import (
	"github.com/CDCgov/phinvads-go/internal/app"
	"github.com/CDCgov/phinvads-go/internal/config"
)

func main() {
	cfg := config.LoadConfig()

	pv := app.SetupApp(cfg)

	pv.Run()
}
