package routes

import (
	profiles "backend-golang/controllers/profiles"
	stockhistory "backend-golang/controllers/stock_history"
	stockins "backend-golang/controllers/stock_ins"
	stockouts "backend-golang/controllers/stock_outs"
	stocks "backend-golang/controllers/stocks"
	users "backend-golang/controllers/users"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	LoggerMiddleware       echo.MiddlewareFunc
	JWTMiddleware          echojwt.Config
	AuthController         users.AuthController
	ProfilesController     profiles.ProfilesController
	StocksController       stocks.StockController
	StockInsController     stockins.StockInController
	StockOutsController    stockouts.StockOutController
	StockHistoryController stockhistory.StockHistoryController
}

func (cl *ControllerList) RegisterRoutes(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)

	auth := e.Group("auth")

	auth.POST("/register", cl.AuthController.Register)
	auth.POST("/login", cl.AuthController.Login)
	auth.POST("/logout", cl.AuthController.Logout)

	users := e.Group("users", echojwt.WithConfig(cl.JWTMiddleware))
	users.GET("/:id", cl.ProfilesController.GetByID)
	users.PUT("/profiles/customer/:id", cl.ProfilesController.UpdateProfileUser)
	users.PUT("/profiles/picture/:id", cl.ProfilesController.UploadProfileImage)

	stocks := e.Group("stocks", echojwt.WithConfig(cl.JWTMiddleware))
	stocks.GET("/:id", cl.StocksController.GetByID)
	stocks.GET("/:id", cl.StocksController.DownloadBarcodeByID)
	stocks.POST("", cl.StocksController.Create)
	// stocks.PUT("/stock_in/:id", cl.StocksController.StockIn)
	// stocks.PUT("/stock_out/:id", cl.StocksController.StockOut)

	stockIns := e.Group("stock_ins", echojwt.WithConfig(cl.JWTMiddleware))
	stockIns.POST("/create", cl.StockInsController.Create)
	stockIns.POST("", cl.StockInsController.StockIn)

	// stockouts := e.Group("stock_outs", echojwt.WithConfig(cl.JWTMiddleware))
	stockouts := e.Group("stock_outs")
	stockouts.GET("", cl.StockOutsController.GetAll)
	stockouts.POST("", cl.StockOutsController.StockOut)
	stockouts.GET("/download", cl.StockOutsController.ExportToExcel)

	stockhistory := e.Group("stock_history", echojwt.WithConfig(cl.JWTMiddleware))
	stockhistory.POST("", cl.StockHistoryController.Create)
}
