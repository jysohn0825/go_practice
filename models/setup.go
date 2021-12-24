package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("mysql", "root:Kona!234@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Book{})

	DB = database
}
