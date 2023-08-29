package database

import (
	"echo-graphql-qiita/graph/model"
	"echo-graphql-qiita/utils"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Init() {
	host := utils.Getenv("DB_HOST")
	user := utils.Getenv("DB_USER")
	password := utils.Getenv("DB_PASSWORD")
	dbname := utils.Getenv("DB_NAME")
	port := utils.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", host, user, password, dbname, port)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&model.Post{})
}
