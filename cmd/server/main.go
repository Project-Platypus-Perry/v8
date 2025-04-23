package main

import (
	"github.com/gagan-gaurav/v8/pkg/logger"

	_ "github.com/gagan-gaurav/v8/docs" // This line is needed for swagger
	"github.com/gagan-gaurav/v8/internal/app"
	"github.com/gagan-gaurav/v8/internal/config"
)

// @title           Base API
// @version         1.0
// @description     A RESTful API service providing user management and health check endpoints
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
// @schemes   http https

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
