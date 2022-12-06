package handler

import (
	"macus/database"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func AddAPIs(db *sqlx.DB, engine *gin.Engine) {
	customersHandler := CustomerHandler{
		CustomerDB: database.PostgresDB{
			DB: db,
		},
	}
	customersV1 := engine.Group("/api/v1")
	{
		customers := customersV1.Group("/customers")
		{
			customers.GET(":id", customersHandler.GetCustomer)
			customers.GET("", customersHandler.ListCustomers)
			customers.POST("", customersHandler.CreateCustomer)
			customers.DELETE(":id", customersHandler.DeleteCustomer)
			customers.PATCH(":id", customersHandler.UpdateCustomer)
		}
	}

	authenHandler := AuthenHandler{}
	authenv1 := engine.Group("/api/v1")
	{
		authen := authenv1.Group("/authen")
		{
			authen.POST("/LogIn", authenHandler.LogIn)
		}
	}

	engine.GET("/healthcheck", HealthCheck)
}

func AddSwagger(engine *gin.Engine) {
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
