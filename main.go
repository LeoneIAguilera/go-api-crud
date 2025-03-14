package main

import (
	"github.com/LeoneIAguilera/web-simple-two/controllers"
	"github.com/LeoneIAguilera/web-simple-two/initializers"
	"github.com/LeoneIAguilera/web-simple-two/internal"
	"github.com/LeoneIAguilera/web-simple-two/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDb()
}

func main() {
	r := gin.Default()

	//User Controllers
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate",middleware.RequireAuth, controllers.Validate)

	// Create routes
	r.POST("/sales/create", middleware.RequireAuth, internal.CreateSales)
	r.POST("/payments/create", middleware.RequireAuth, internal.CreatePayments)
	r.POST("/debts/create", middleware.RequireAuth, internal.CreateDebt)
	r.POST("/suppliers/create", middleware.RequireAuth, internal.CreateSupplier)

	// Delete routes
	r.DELETE("/sales/:id", middleware.RequireAuth, internal.DeleteSales)
	r.DELETE("/payments/:id", middleware.RequireAuth, internal.DeletePayments)
	r.DELETE("/debts/:id", middleware.RequireAuth, internal.DeleteDebt)
	r.DELETE("/suppliers/:id", middleware.RequireAuth, internal.DeleteSupplier)

	// Updates routes
	r.PUT("/payments/:id", middleware.RequireAuth, internal.UpdatePayments)
	r.PUT("/debts/:id", middleware.RequireAuth, internal.UpdateDebt)
	r.PUT("/suppliers/:id", middleware.RequireAuth, internal.UpdateSupplier)

	// RetrieveAll
	r.GET("/sales", middleware.RequireAuth, internal.RetrieveAllSales)
	r.GET("/payments", middleware.RequireAuth, internal.RetrieveAllPayments)
	r.GET("/debts", middleware.RequireAuth, internal.RetrieveAllDebt)
	r.GET("/suppliers", middleware.RequireAuth, internal.RetrieveAllSupplier)

	r.Run()
}