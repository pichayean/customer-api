package main

import (
	"log"

	"macus/config"
	"macus/database"
	_ "macus/docs"
	"macus/handler"

	"github.com/gin-gonic/gin"
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
	db, err := database.Open(config.GetDBConfig())
	if err != nil {
		log.Fatal("connecting database fail", err)
	}
	engine := gin.Default()
	handler.AddAPIs(db, engine)
	handler.AddSwagger(engine)
	log.Fatal((engine.Run(":8080")))
}

// https://dev.to/billylkc/parse-json-api-response-in-go-10ng
