package internal

import (
	"net/http"
	"strconv"

	"github.com/LeoneIAguilera/web-simple-two/initializers"
	"github.com/LeoneIAguilera/web-simple-two/models"
	"github.com/gin-gonic/gin"
)

func updateEntity(c *gin.Context, entity interface{}, upEntity interface{}, entityName string) {
	id := c.Param("id")
	entityID, err := strconv.Atoi(id)
	
	if err != nil || entityID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to get id",
		})
		return
	}
	
	result := initializers.DB.Model(entity).Where("id = ?", entityID).Updates(upEntity)
	
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Error updating " + entityName,
		})

		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": entityName + "Not found",
		})

		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": entityName + " updated successfully!",
	})
}

func checkBody(c *gin.Context, body interface{}) {
	if c.Bind(body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read body",
		})

		return
	}
}

func UpdatePayments(c *gin.Context) {
	var body struct {
		Amount 		float64
		SupplierID  uint
	}
	checkBody(c, &body)
	
	var payments models.Payments
	update := models.Payments{
		Amount: body.Amount,
		SupplierID: body.SupplierID,
	}
	
	
	updateEntity(c, &payments, update, "payments")
}

func UpdateSupplier(c *gin.Context) {
	var body struct {
		Name 	string
	}
	checkBody(c, &body)
	
	var supplier models.Supplier
	update := models.Supplier{
		Name: body.Name,
	}
	
	
	updateEntity(c, &supplier, update, "debt")
}
func UpdateDebt(c *gin.Context) {
	var body struct {
		Amount 			float64
		Description 	string
		SupplierID  	uint
	}
	checkBody(c, &body)
	
	var debt models.Debt
	update := models.Debt{
		Amount: body.Amount,
		Description: body.Description,
		SupplierID: body.SupplierID,
	}
	updateEntity(c, &debt, update, "debt")
	
}