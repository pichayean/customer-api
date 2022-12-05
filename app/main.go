package main

import (
	"flag"
	"log"

	"macus/database"
	_ "macus/docs"
	"macus/handler"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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
	var port, dbhost, dbschema, dbusername, dbpassword, disableTLS string
	var dbport int
	flag.StringVar(&port, "port", "8080", "port for open service")
	flag.StringVar(&dbhost, "dbhost", "144.126.140.118", "database host name")
	flag.IntVar(&dbport, "dbport", 5432, "database port")
	flag.StringVar(&dbschema, "dbschema", "customers", "database schema name")
	flag.StringVar(&dbusername, "dbusername", "postgres", "database user name")
	flag.StringVar(&dbpassword, "dbpassword", "xxx", "database password")
	flag.StringVar(&disableTLS, "disableTLS", "Y", "database disableTLS[Y/n]")
	flag.Parse()
	var databaseTSL bool
	if disableTLS == "n" {
		databaseTSL = false
	} else {
		databaseTSL = true
	}
	dbConfig := database.Config{
		User:       dbusername,
		Password:   dbpassword,
		Host:       dbhost,
		Port:       dbport,
		Name:       dbschema,
		DisableTLS: databaseTSL,
	}

	db, err := database.Open(dbConfig)
	if err != nil {
		log.Fatal("connecting database fail", err)
	}
	r := gin.Default()

	// r := gin.New()

	r.GET("/healthcheck", handler.HealthCheckHandler)
	// r.Use(middlewares.AuthUser())

	API(db, r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Fatal((r.Run(":8080")))
}

func API(db *sqlx.DB, engine *gin.Engine) {

	customersHandler := handler.CustomerHandler{
		CustomerDB: database.PostgresDB{
			DB: db,
		},
	}
	v1 := engine.Group("/api/v1")
	{
		customers := v1.Group("/customers")
		{
			// g.POST("/image", middleware.AuthUser(h.TokenService), h.Image)
			customers.GET(":id", customersHandler.GetCustomer)
			customers.GET("", customersHandler.ListCustomers)
			customers.POST("", customersHandler.CreateCustomer)
			customers.DELETE(":id", customersHandler.DeleteCustomer)
			customers.PATCH(":id", customersHandler.UpdateCustomer)
		}
	}
}

// https://dev.to/billylkc/parse-json-api-response-in-go-10ng
