package main

import (
	"github.com/CDCgov/phinvads-fhir/internal/app"
	"github.com/CDCgov/phinvads-fhir/internal/config"
)

func main() {
	cfg := config.LoadConfig()

	pv := app.SetupApp(cfg)

	pv.Run()
}
