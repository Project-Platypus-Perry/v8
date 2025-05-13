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
	BatchService        service.BatchService
	ClassroomService    service.ClassroomService
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
	batchHandler := handler.NewBatchHandler(r.deps.BatchService)
	// classroomHandler := handler.NewClassroomHandler(r.deps.ClassroomService)

	// Admin only routes
	adminRoutes := protected.Group("")
	adminRoutes.Use(middleware.RequireRole(constants.AdminRole))
	adminRoutes.POST("/users/invite", userHandler.InviteUsers)
	adminRoutes.DELETE("/users/:id", userHandler.DeleteUser)
	adminRoutes.POST("/batch", batchHandler.CreateBatch)
	adminRoutes.POST("/batch/users/add", batchHandler.AddUserToBatch)
	adminRoutes.POST("/batch/users/remove", batchHandler.RemoveUserFromBatch)
	// adminRoutes.GET("/batch/:id", userHandler.GetBatch)
	// adminRoutes.PATCH("/batch/:id", userHandler.UpdateBatch)
	// adminRoutes.DELETE("/batch/:id", userHandler.DeleteBatch)

	// Admin and Instructor routes
	staffRoutes := protected.Group("")
	staffRoutes.Use(middleware.RequireRole(constants.AdminRole, constants.InstructorRole))
	staffRoutes.PATCH("/users/:id", userHandler.UpdateUser)
	// staffRoutes.POST("/classroom", userHandler.CreateClassroom)
	// staffRoutes.PATCH("/classroom/:id", userHandler.UpdateClassroom)
	// staffRoutes.DELETE("/classroom/:id", userHandler.DeleteClassroom)

	// add user(student) to classroom
	// staffRoutes.POST("/users/:id/classroom", userHandler.AddUserToClassroom)
	// remove user(student) from classroom
	// staffRoutes.DELETE("/users/:id/classroom/:classroomId", userHandler.RemoveUserFromClassroom)

	// All authenticated users can read
	protected.GET("/users/:id", userHandler.GetUser)
	protected.POST("/users/request-reset-password", userHandler.RequestPasswordReset)
	protected.POST("/users/reset-password", userHandler.ResetPassword)
	// protected.GET("/classroom", userHandler.ListClassrooms)
	// protected.GET("/classroom/:id", userHandler.GetClassroom)
	protected.GET("/batch/list", batchHandler.ListUserBatches)
	protected.GET("/batch/:id", batchHandler.GetBatch)
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
			if err := c.NoContent(code); err != nil {
				logger.Error("Error in customHTTPErrorHandler", zap.Error(err))
			}
		} else {
			var responseErr error
			if err == echo.ErrNotFound {
				responseErr = response.Error(c, code, "Route not found")
			} else {
				responseErr = response.Error(c, code, message)
			}
			if responseErr != nil {
				logger.Error("Error in customHTTPErrorHandler", zap.Error(responseErr))
			}
		}
	}
}
