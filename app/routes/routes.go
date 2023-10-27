package routes

import (
	"backend-golang/controllers/users"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	LoggerMiddleware echo.MiddlewareFunc
	JWTMiddleware    echojwt.Config
	AuthController   users.AuthController
}

func (cl *ControllerList) RegisterRoutes(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)

	auth := e.Group("auth")

	auth.POST("/register", cl.AuthController.Register)
	auth.POST("/login", cl.AuthController.Login)
	auth.POST("/logout", cl.AuthController.Logout)

	users := e.Group("users", echojwt.WithConfig(cl.JWTMiddleware))
	users.PUT("/profiles/customer/:id", cl.AuthController.UpdateProfileUser)
	users.PUT("/profiles/picture/:id", cl.AuthController.UploadProfileImage)

	// course := e.Group("/api/v1/courses", echojwt.WithConfig(cl.JWTMiddleware))
	// course.Use(middlewares.VerifyToken)

	// course.GET("", cl.CourseController.GetAll)
	// course.GET("/:id", cl.CourseController.GetByID)
	// course.POST("", cl.CourseController.Create)
	// course.PUT("/:id", cl.CourseController.Update)
	// course.DELETE("/:id", cl.CourseController.Delete)
	// course.POST("/:id", cl.CourseController.Restore)
	// course.DELETE("/:id/force", cl.CourseController.ForceDelete)

	// category := e.Group("/api/v1/categories", echojwt.WithConfig(cl.JWTMiddleware))
	// category.Use(middlewares.VerifyToken)

	// category.GET("", cl.CategoryController.GetAll)
	// category.POST("", cl.CategoryController.Create)
	// category.PUT("/:id", cl.CategoryController.Update)
	// category.DELETE("/:id", cl.CategoryController.Delete)
}
