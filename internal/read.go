package internal

import (
	"net/http"

	"github.com/LeoneIAguilera/web-simple-two/initializers"
	"github.com/LeoneIAguilera/web-simple-two/models"
	"github.com/gin-gonic/gin"
)

func retrieveAllEntityPreload(c *gin.Context, entity interface{}, entityName string, preLoad string) {
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

func retrieveAllEntity(c *gin.Context, entity interface{}, entityName string) {
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

func RetrieveAllSales(c *gin.Context) {
	var sales []models.Sales
	retrieveAllEntity(c, &sales, "sales")
}

func RetrieveAllPayments(c *gin.Context) {
	var payments []models.Payments
	retrieveAllEntityPreload(c, &payments, "payments", "Supplier")
}
func RetrieveAllDebt(c *gin.Context) {
	var debt []models.Debt
	retrieveAllEntityPreload(c, &debt, "debt", "Supplier")
}
func RetrieveAllSupplier(c *gin.Context) {
	var supplier []models.Supplier
	retrieveAllEntity(c, &supplier, "supplier")
}