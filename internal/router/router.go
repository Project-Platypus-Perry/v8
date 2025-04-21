package router

import (
	"github.com/gagan-gaurav/base/internal/config"
	"github.com/gagan-gaurav/base/internal/handler"
	"github.com/gagan-gaurav/base/internal/middleware"
	"github.com/gagan-gaurav/base/internal/service"
	"github.com/labstack/echo/v4"
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
	v1 := r.e.Group("/api/v1")
	v1.Use(middleware.RequestLogger)
	// v1.Use(middleware.JWTAuth)

	// Health check routes
	health := handler.NewHealthHandler()
	v1.GET("/health", health.Check)

	// User routes
	user := handler.NewUserHandler(r.deps.UserService)
	v1.GET("/users", user.ListUsers)
}
