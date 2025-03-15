package main

import (
	"time"

	"github.com/LeoneIAguilera/web-simple-two/controllers"
	"github.com/LeoneIAguilera/web-simple-two/initializers"
	"github.com/LeoneIAguilera/web-simple-two/internal"
	"github.com/LeoneIAguilera/web-simple-two/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDb()
}

func main() {
	//close DB
	conn, err := initializers.DB.DB()
	if err != nil {
		panic("Failed to get instance")
	}

	defer conn.Close()

	r := gin.Default()
	
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Permite solicitudes desde el frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}, // Métodos permitidos
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Encabezados permitidos
		ExposeHeaders:    []string{"Content-Length"}, // Encabezados expuestos
		AllowCredentials: true, // Permite el envío de cookies
		MaxAge:           12 * time.Hour, // Duración del preflight cache
	}))
	api := r.Group("/api")
	
	//User Controllers
	api.POST("/signup", controllers.Signup)
	api.POST("/login", controllers.Login)
	api.GET("/validate",middleware.RequireAuth, controllers.Validate)
	api.POST("/logout",middleware.RequireAuth, controllers.Logout)
	
	// Create routes
	api.POST("/sales/create", middleware.RequireAuth, internal.CreateSales)
	api.POST("/payments/create", middleware.RequireAuth, internal.CreatePayments)
	api.POST("/debts/create", middleware.RequireAuth, internal.CreateDebt)
	api.POST("/suppliers/create", middleware.RequireAuth, internal.CreateSupplier)

	// Delete routes
	api.DELETE("/sales/:id", middleware.RequireAuth, internal.DeleteSales)
	api.DELETE("/payments/:id", middleware.RequireAuth, internal.DeletePayments)
	api.DELETE("/debts/:id", middleware.RequireAuth, internal.DeleteDebt)
	api.DELETE("/suppliers/:id", middleware.RequireAuth, internal.DeleteSupplier)

	// Updates routes
	api.PUT("/payments/:id", middleware.RequireAuth, internal.UpdatePayments)
	api.PUT("/debts/:id", middleware.RequireAuth, internal.UpdateDebt)
	api.PUT("/suppliers/:id", middleware.RequireAuth, internal.UpdateSupplier)

	// RetrieveAll
	api.GET("/sales", middleware.RequireAuth, internal.ViewSales)
	api.GET("/payments", middleware.RequireAuth, internal.ViewPayments)
	api.GET("/debts", middleware.RequireAuth, internal.ViewDebts)
	api.GET("/suppliers", middleware.RequireAuth, internal.ViewSuppliers)

	r.Run(":8080")
}