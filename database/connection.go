package database

import (
	"github.com/sergio-rey/jwt/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	connection, err := gorm.Open(mysql.Open("root:@/jwt?parseTime=true"), &gorm.Config{})

	if err != nil {
		panic("Could not connect with database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
