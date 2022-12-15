package main

import (
	"log"

	"macus/database"
	_ "macus/docs"
	"macus/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title Customers API
// @version 1.0
// @description.markdown
// @termsOfService http://somewhere.com/

// @contact.name API Support
// @contact.url http://somewhere.com/support
// @contact.email support@somewhere.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @schemes https http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// db, err := database.Open(config.GetDBConfig())
	// if err != nil {
	// 	log.Fatal("connecting database fail", err)
	// }
	dsn := "host=144.126.140.118 user=postgres password=Ld4t5555 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// db, err := gorm.Open(postgres.Open(os.Getenv("DB_CONN")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&database.CustomerEntity{})
	engine := gin.Default()
	gormStore := database.NewGormStore(db)
	handlers.AddAPIs(gormStore, engine)
	handlers.AddSwagger(engine)
	log.Fatal((engine.Run(":8080")))
}

// https://dev.to/billylkc/parse-json-api-response-in-go-10ng
