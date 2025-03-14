package internal

import (
	"net/http"
	"strconv"

	"github.com/LeoneIAguilera/web-simple-two/initializers"
	"github.com/LeoneIAguilera/web-simple-two/models"
	"github.com/gin-gonic/gin"
)

func deleteEntity(c *gin.Context, entity interface{}, entityName string) {
	id := c.Param("id")

	entityID, err := strconv.Atoi(id)

	if err != nil {		
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to get id",
		})

		return
	}
	
	checkID := initializers.DB.Where("id = ?", entityID).First(&entity)

	if checkID.Error != nil || checkID.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": entityName + " not found!",
		})

		return
	}

	result := initializers.DB.Where("id = ?", entityID).Delete(&entity)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Failed to delete" + entityName,
		})

		return
	}

	c.Status(http.StatusNoContent)
}


func DeleteSales(c *gin.Context) {
	var sales models.Sales
	deleteEntity(c, &sales, "sales")
}

func DeletePayments(c *gin.Context) {
	var payments models.Payments
	deleteEntity(c, &payments, "payments")
}

func DeleteDebt(c *gin.Context) {
	var debt models.Debt
	deleteEntity(c, &debt, "debt")
}

func DeleteSupplier(c *gin.Context) {
	var supplier models.Supplier
	deleteEntity(c, &supplier, "supplier")
}

