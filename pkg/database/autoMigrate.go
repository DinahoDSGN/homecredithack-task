package database

import (
	"gorm.io/gorm"
	"homecredithack-task/pkg/models"
	"log"
)

func InitMigrate(connection *gorm.DB) error {
	Drop(connection)

	Push(connection)

	Mocks(connection)

	return nil
}

func Push(connection *gorm.DB) error {
	connection.AutoMigrate(&models.CreditCard{})
	connection.AutoMigrate(&models.Shop{})
	connection.AutoMigrate(&models.Category{})
	connection.AutoMigrate(&models.Product{})
	connection.AutoMigrate(&models.Product_Category{})
	connection.AutoMigrate(&models.Transaction{})
	connection.AutoMigrate(&models.Order{})
	connection.AutoMigrate(&models.Order_item{})
	connection.AutoMigrate(&models.CreditCard_item{})

	log.Println("Models successfully migrated")

	return nil
}

func Drop(connection *gorm.DB) error {
	connection.Migrator().DropTable("shops")
	connection.Migrator().DropTable("credit_cards")
	connection.Migrator().DropTable("products")
	connection.Migrator().DropTable("categories")
	connection.Migrator().DropTable("product_categories")
	connection.Migrator().DropTable("transactions")
	connection.Migrator().DropTable("orders")
	connection.Migrator().DropTable("order_items")
	connection.Migrator().DropTable("creditcard_items")

	return nil
}

func Mocks(connection *gorm.DB) error {
	query := models.CreditCard{
		Number: "666-666-666",
		Money:  666,
	}
	query2 := models.CreditCard{
		Number: "asd456dsa456",
		Money:  468,
	}
	query3 := models.Shop{
		Name: "Aknur",
	}
	query4 := models.Product{
		Name:        "VIVO",
		Description: "Notebook Ultra",
		Price:       220000,
	}
	query5 := models.Category{
		Title: "Notebook",
	}
	query6 := models.Product_Category{
		ProductId:  1,
		CategoryId: 1,
	}
	query7 := models.Order{
		Status: true,
		Amount: 2,
		Mobile: "+77479797932",
	}
	query8 := models.Order_item{
		ProductId: 1,
		OrderId:   1,
	}
	query9 := models.Transaction{
		OrderId: 1,
		Type:    "",
		Code:    "",
		Status:  true,
	}

	connection.Create(&query)
	connection.Create(&query2)
	connection.Create(&query3)
	connection.Create(&query4)
	connection.Create(&query5)
	connection.Create(&query6)
	connection.Create(&query7)
	connection.Create(&query8)
	connection.Create(&query9)

	return nil
}
