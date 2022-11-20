package config

import (
	"assigment-2/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func StartDB() {
	var err error
	dsn := "root:root@tcp(localhost)/rest-api-assigment-2?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		QueryFields:            true,
	})
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&models.Item{})
	DB.AutoMigrate(&models.Order{})

}

func GetDB() *gorm.DB {
	return DB
}

// heroes_db
