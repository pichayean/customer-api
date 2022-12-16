package main

import (
	"log"
	"os"

	"macus/data"
	"macus/entities"
	"macus/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load(os.Getenv("APP_ENV") + ".env")
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_CONN")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&entities.CustomerEntity{})
	engine := gin.Default()
	gormStore := data.NewGormProvider(db)
	router.AddAPIs(gormStore, engine)
	log.Fatal((engine.Run(":" + os.Getenv("PORT"))))
}

// https://dev.to/billylkc/parse-json-api-response-in-go-10ng
