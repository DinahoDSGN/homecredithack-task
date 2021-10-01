package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	connection, err := gorm.Open(mysql.Open("root:root@/homecredithack"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	err = InitMigrate(DB)
	if err != nil {
		log.Fatal("Migrations failed!")
	}

	return DB
}
