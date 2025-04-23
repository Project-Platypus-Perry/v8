package app

// Initialize the app, dependencies, and router
import (
	"context"
	"fmt"

	"github.com/gagan-gaurav/v8/internal/config"
	"github.com/gagan-gaurav/v8/internal/db"
	"github.com/gagan-gaurav/v8/internal/repository"
	"github.com/gagan-gaurav/v8/internal/router"
	"github.com/gagan-gaurav/v8/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// CustomValidator is a wrapper for the validator
type CustomValidator struct {
	validator *validator.Validate
}

// Validate implements echo.Validator interface
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

type App struct {
	cfg  *config.Config
	e    *echo.Echo
	deps *Dependencies
}

type Dependencies struct {
	UserService service.UserService
}

func NewDependencies(db *gorm.DB, cfg *config.Config) *Dependencies {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	return &Dependencies{
		UserService: userService,
	}
}

func NewApp(cfg *config.Config) *App {
	// Initialize the database
	db := db.InitPostgres(cfg)

	// Initialize the dependencies
	deps := NewDependencies(db, cfg)

	// Initialize Echo
	e := echo.New()

	// Register validator
	e.Validator = &CustomValidator{validator: validator.New()}

	// Initialize router
	routerDeps := &router.Dependencies{
		UserService: deps.UserService,
	}
	r := router.NewRouter(e, cfg, routerDeps)
	r.InitRoutes()

	return &App{
		cfg:  cfg,
		e:    e,
		deps: deps,
	}
}

func (a *App) Start() error {
	return a.e.Start(":" + a.cfg.Port)
}

func (a *App) Stop() error {
	return a.e.Shutdown(context.Background())
}

func (a *App) Router() *echo.Echo {
	fmt.Println("Router")
	return a.e
}

func (a *App) Dependencies() *Dependencies {
	return a.deps
}
