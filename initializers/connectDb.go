package initializers

import (
	"log"
	"os"

	"github.com/LeoneIAguilera/web-simple-two/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	
	dsn := os.Getenv("DB_URL")
	
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		panic("Failed to connect database")
	}
	
	migrateModels := []any{
		&models.User{},
		&models.Debt{},
		&models.Sales{},
		&models.Supplier{},
		&models.Payments{},
	}

	for _, model := range migrateModels {
		if err := DB.AutoMigrate(model); err != nil {
			log.Fatalf("Error to migrate the model: %T error: %v", model, err)
		}
		
	}
}