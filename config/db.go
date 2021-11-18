package config

import (
	"fmt"
	"jwt-gin/models"
	"jwt-gin/utils"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	var username string = "root"
	var password string = "bds29"
	var host string = "127.0.0.1"
	var port string = "3306"
	var database string = "database_movie"

	environment := utils.Getenv("ENVIRONMENT", "development")

	if environment == "production" {
		username = os.Getenv("DATABASE_USERNAME")
		password = os.Getenv("DATABASE_PASSWORD")
		host = os.Getenv("DATABASE_HOST")
		port = os.Getenv("DATABASE_PORT")
		database = os.Getenv("DATABASE_NAME")

		dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " sslmode=require"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err.Error())
		}

		db.AutoMigrate(&models.User{}, &models.Movie{}, &models.AgeRatingCategory{})
		return db
	} else {
		dsn := fmt.Sprintf(
			"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
			username, password, host, port, database)
		fmt.Println(dsn)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			panic(err.Error())
		}

		db.AutoMigrate(&models.Movie{}, &models.AgeRatingCategory{}, &models.User{})
		return db
	}
}
