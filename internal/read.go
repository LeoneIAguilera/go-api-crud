package internal

import (
	"net/http"

	"github.com/LeoneIAguilera/web-simple-two/initializers"
	"github.com/LeoneIAguilera/web-simple-two/models"
	"github.com/gin-gonic/gin"
)

func viewAllEntityPreload(c *gin.Context, entity interface{}, entityName string, preLoad string) {
	result := initializers.DB.Preload(preLoad).Find(entity)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": result.Error.Error(),
		})

		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "No " + entityName + "found",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		entityName: entity,
	})
}

func viewAllEntity(c *gin.Context, entity interface{}, entityName string) {
	result := initializers.DB.Find(entity)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": result.Error.Error(),
		})

		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "No " + entityName + "found",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		entityName: entity,
	})
}

func ViewSales(c *gin.Context) {
	var sales []models.Sales
	viewAllEntity(c, &sales, "sales")
}

func ViewPayments(c *gin.Context) {
	var payments []models.Payments
	viewAllEntityPreload(c, &payments, "payments", "Supplier")
}
func ViewDebts(c *gin.Context) {
	var debt []models.Debt
	viewAllEntityPreload(c, &debt, "debts", "Supplier")
}
func ViewSuppliers(c *gin.Context) {
	var supplier []models.Supplier
	viewAllEntity(c, &supplier, "suppliers")
}