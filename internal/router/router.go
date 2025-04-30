package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/project-platypus-perry/v8/internal/config"
	"github.com/project-platypus-perry/v8/internal/constants"
	"github.com/project-platypus-perry/v8/internal/handler"
	"github.com/project-platypus-perry/v8/internal/middleware"
	"github.com/project-platypus-perry/v8/internal/service"
	"github.com/project-platypus-perry/v8/pkg/logger"
	"github.com/project-platypus-perry/v8/pkg/response"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/zap"
)

type Router struct {
	e    *echo.Echo
	cfg  *config.Config
	deps *Dependencies
}

type Dependencies struct {
	UserService         service.UserService
	AuthService         service.AuthService
	OrganizationService service.OrganizationService
}

func NewRouter(e *echo.Echo, cfg *config.Config, deps *Dependencies) *Router {
	return &Router{
		e:    e,
		cfg:  cfg,
		deps: deps,
	}
}

func (r *Router) InitRoutes() {
	// Custom HTTP error handler
	r.e.HTTPErrorHandler = customHTTPErrorHandler

	// Initialize middlewares
	rateLimiter := middleware.NewRateLimiter(r.cfg.RateLimiter)
	jwtMiddleware := middleware.NewJWTMiddleware(r.cfg.JWT)

	// Apply global middlewares
	r.e.Use(middleware.RequestLogger)
	r.e.Use(rateLimiter.RateLimit)

	// Add not found handler for undefined routes
	r.e.Any("/*", handleNotFound)

	// API v1 group
	v1 := r.e.Group("/api/v1")

	// Public routes
	v1.GET("/swagger/*", echoSwagger.WrapHandler)
	v1.GET("/health", handler.NewHealthHandler().Check)

	// Auth routes (no JWT required)
	authHandler := handler.NewAuthHandler(r.deps.AuthService)
	auth := v1.Group("/auth")
	auth.POST("/refresh", jwtMiddleware.RefreshToken)
	auth.POST("/register", authHandler.RegisterOrganization)
	auth.POST("/login", authHandler.Login)

	// Protected routes
	protected := v1.Group("")
	protected.Use(jwtMiddleware.JWTAuth)

	// User routes with RBAC
	userHandler := handler.NewUserHandler(r.deps.UserService)

	// Admin only routes
	adminRoutes := protected.Group("")
	adminRoutes.Use(middleware.RequireRole(constants.AdminRole))
	adminRoutes.POST("/users/invite", userHandler.InviteUsers)
	adminRoutes.DELETE("/users/:id", userHandler.DeleteUser)

	// Admin and Instructor routes
	staffRoutes := protected.Group("")
	staffRoutes.Use(middleware.RequireRole(constants.AdminRole, constants.InstructorRole, constants.StudentRole))
	staffRoutes.PATCH("/users/:id", userHandler.UpdateUser)

	// All authenticated users can read
	protected.GET("/users/:id", userHandler.GetUser, middleware.RequirePermission(constants.ReadUser))
	protected.POST("/users/request-reset-password", userHandler.RequestPasswordReset)
	protected.POST("/users/reset-password", userHandler.ResetPassword)
}

// handleNotFound handles undefined routes
func handleNotFound(c echo.Context) error {
	return response.NotFound(c, "Route not found")
}

// customHTTPErrorHandler handles all HTTP errors
func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := "Internal Server Error"

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		if msg, ok := he.Message.(string); ok {
			message = msg
		} else if msg, ok := he.Message.(error); ok {
			message = msg.Error()
		} else {
			message = "An error occurred"
		}
	}

	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(code)
			logger.Error("Error in customHTTPErrorHandler", zap.Error(err))
		} else {
			err = response.Error(c, code, message)
			logger.Error("Error in customHTTPErrorHandler", zap.Error(err))
		}
	}
}
