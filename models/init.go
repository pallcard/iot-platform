package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func NewDB() {
	dsn := "root:lk961232@tcp(192.168.101.50:3308)/iot_platform?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal("[DB ERROR]:", err)
	}
	err = db.AutoMigrate(&UserBasic{}, &ProductBasic{}, &DeviceBasic{})
	if err != nil {
		log.Fatal("[DB ERROR]:", err)
	}
	DB = db
}
