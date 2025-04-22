package main

import (
	"github.com/gagan-gaurav/base/pkg/logger"

	"github.com/gagan-gaurav/base/internal/app"
	"github.com/gagan-gaurav/base/internal/config"
)

// @title           My Enterprise API
// @version         1.0
// @description     This is a sample enterprise backend API.
// @termsOfService  http://yourdomain.com/terms/

// @contact.name   API Support
// @contact.email  support@yourdomain.com

// @host      localhost:8080
// @BasePath  /api/v1

func main() {
	cfg := config.Load()
	logger := logger.Init(cfg.LogLevel)

	// Initialize the app
	logger.Info("Initializing the app")
	app := app.NewApp(cfg)
	logger.Info("App initialized")

	// Start the app
	logger.Info("Starting the app")
	err := app.Start()
	if err != nil {
		logger.Error(err.Error())
	}
	// logger.Info("App started")

	// // Initialize the router for the app
	// logger.Info("Initializing the router")
	// router.InitRoutes(app.Router(), cfg, app.Dependencies())
	// logger.Info("Router initialized")
}
