package main

import (
	"jwt-gin/config"
	"jwt-gin/routes"
	"jwt-gin/utils"
	"log"

	"github.com/joho/godotenv"
	"github.com/swaggo/swag/example/basic/docs"
)

func main() {
	// for load godotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Movie."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	portApi := utils.Getenv("PORT", ":7877")
	docs.SwaggerInfo.Host = utils.Getenv("SWAGGER_HOST", ":9090")

	// database connection
	db := config.ConnectDatabase()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// router
	r := routes.SetupRouter(db)
	r.Run(portApi)
}
