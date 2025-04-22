package router

import (
	"net/http"

	"github.com/gagan-gaurav/base/internal/config"
	"github.com/gagan-gaurav/base/internal/handler"
	"github.com/gagan-gaurav/base/internal/middleware"
	"github.com/gagan-gaurav/base/internal/service"
	"github.com/gagan-gaurav/base/pkg/logger"
	"github.com/labstack/echo/v4"
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
