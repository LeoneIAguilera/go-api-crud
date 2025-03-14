package internal

import (
	"net/http"

	"github.com/LeoneIAguilera/web-simple-two/initializers"
	"github.com/LeoneIAguilera/web-simple-two/models"
	"github.com/gin-gonic/gin"
)

func createEntity(c *gin.Context, entity any, entityName string) {

	result := initializers.DB.Create(entity)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to create a new " + entityName,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully!",
	})
}

func CreateSales(c *gin.Context) {
	var body struct {
		Amount	float64
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read body",
		})
		return
	}

	newSale := models.Sales{
		Amount: body.Amount,
	}
	createEntity(c, &newSale, "sales")
}

func CreatePayments(c *gin.Context) {
	var body struct {
		Amount	float64
		SupplierID uint
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read body",
		})
		return
	}

	newPayment := models.Payments{
		Amount: body.Amount,
		SupplierID: body.SupplierID,
	}
	createEntity(c, &newPayment, "payment")
}

func CreateDebt(c *gin.Context) {
	var body struct {
		Amount	float64
		Description string
		SupplierID uint
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read body",
		})
		return
	}

	newDebt := models.Debt{
		Amount: body.Amount,
		Description: body.Description,
		SupplierID: body.SupplierID,
	}
	createEntity(c, &newDebt, "debt")
}

func CreateSupplier(c *gin.Context) {
	var body struct {
		Name string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read body",
		})
		return
	}

	newSupplier := models.Supplier{
		Name: body.Name,
	}
	createEntity(c, &newSupplier, "supplier")
}