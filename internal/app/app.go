package app

// Initialize the app, dependencies, and router
import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/project-platypus-perry/v8/internal/config"
	"github.com/project-platypus-perry/v8/internal/db"
	"github.com/project-platypus-perry/v8/internal/repository"
	"github.com/project-platypus-perry/v8/internal/router"
	"github.com/project-platypus-perry/v8/internal/service"
	emailService "github.com/project-platypus-perry/v8/pkg/email_service"
	"github.com/project-platypus-perry/v8/pkg/logger"
	"go.uber.org/zap"
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
	UserService         service.UserService
	AuthService         service.AuthService
	OrganizationService service.OrganizationService
}

func NewDependencies(db *gorm.DB, cfg *config.Config) *Dependencies {

	// Initialize the repositories
	userRepository := repository.NewUserRepository(db)
	organizationRepository := repository.NewOrganizationRepository(db)

	// Initialize the services
	emailService := emailService.NewEmailService(cfg.Email)
	userService := service.NewUserService(userRepository, emailService, cfg.JWT.AccessTokenSecret)
	organizationService := service.NewOrganizationService(organizationRepository)
	authService := service.NewAuthService(userService, organizationService, cfg.JWT)

	return &Dependencies{
		UserService:         userService,
		AuthService:         authService,
		OrganizationService: organizationService,
	}
}

func NewApp(cfg *config.Config) *App {
	// Initialize the database
	db, err := db.InitPostgres(cfg)
	if err != nil {
		logger.Fatal("Failed to initialize database", zap.Error(err))
	}

	// Initialize the dependencies
	deps := NewDependencies(db, cfg)

	// Initialize Echo
	e := echo.New()

	// Register validator
	e.Validator = &CustomValidator{validator: validator.New()}

	// Initialize router
	routerDeps := &router.Dependencies{
		UserService:         deps.UserService,
		AuthService:         deps.AuthService,
		OrganizationService: deps.OrganizationService,
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
