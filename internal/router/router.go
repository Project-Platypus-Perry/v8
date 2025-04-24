package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/project-platypus-perry/v8/internal/config"
	"github.com/project-platypus-perry/v8/internal/handler"
	"github.com/project-platypus-perry/v8/internal/middleware"
	"github.com/project-platypus-perry/v8/internal/service"
	"github.com/project-platypus-perry/v8/pkg/logger"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/zap"
)

type Router struct {
	e    *echo.Echo
	cfg  *config.Config
	deps *Dependencies
}

type Dependencies struct {
	UserService service.UserService
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

	// Add not found handler for undefined routes
	r.e.Any("/*", handleNotFound)

	v1 := r.e.Group("/api/v1")

	v1.Use(middleware.RequestLogger)
	// v1.Use(middleware.JWTAuth)

	// Swagger documentation endpoint
	v1.GET("/swagger/*", echoSwagger.WrapHandler)

	// Health check routes
	health := handler.NewHealthHandler()
	v1.GET("/health", health.Check)

	// User routes
	user := handler.NewUserHandler(r.deps.UserService)
	v1.POST("/users", user.CreateUser)
	v1.GET("/users/:id", user.GetUser)
	v1.PATCH("/users/:id", user.UpdateUser)
	v1.DELETE("/users/:id", user.DeleteUser)
}

// handleNotFound handles undefined routes
func handleNotFound(c echo.Context) error {
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"error": "Route not found",
		"path":  c.Request().URL.Path,
	})
}

// customHTTPErrorHandler handles all HTTP errors
func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := "Internal Server Error"

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = he.Message.(string)
	}

	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(code)
			logger.Error("Error in customHTTPErrorHandler", zap.Error(err))
		} else {
			err = c.JSON(code, map[string]interface{}{
				"error": message,
				"path":  c.Request().URL.Path,
			})
			logger.Error("Error in customHTTPErrorHandler", zap.Error(err))
		}
	}
}
