package router

import (
	"fmt"
	"io/ioutil"
	"macus/data"
	"macus/entities"
	"macus/handlers"
	"macus/middlewares"
	"macus/services"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AddAPIs(db *data.GormProvider[entities.CustomerEntity], engine *gin.Engine) {
	customersHandler := handlers.NewCustomerHandler(db)
	customersV1 := engine.Group("/api/v1")
	{
		customers := customersV1.Group("/customers")
		{
			customers.GET(":id", middlewares.AuthUser(), customersHandler.GetCustomer)
			customers.GET("", customersHandler.ListCustomers)
			customers.POST("", customersHandler.CreateCustomer)
			customers.DELETE(":id", customersHandler.DeleteCustomer)
			customers.PATCH(":id", customersHandler.UpdateCustomer)
		}
	}
	// load rsa keys
	privKeyFile := os.Getenv("PRIV_KEY_FILE")
	fmt.Println(privKeyFile)
	priv, err := ioutil.ReadFile(privKeyFile)
	if err != nil {
		fmt.Println("could not read private key pem file: %w", err)
	}
	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(priv)
	fmt.Println(privKey)
	if err != nil {
		fmt.Println("could not parse private key: %w", err)
	}

	pubKeyFile := os.Getenv("PUB_KEY_FILE")
	fmt.Println(pubKeyFile)
	pub, err := ioutil.ReadFile(pubKeyFile)
	if err != nil {
		fmt.Println("could not read public key pem file: %w", err)
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pub)
	if err != nil {
		fmt.Println("could not parse public key: %w", err)
	}

	// load expiration lengths from env variables and parse as int
	idTokenExp := os.Getenv("ID_TOKEN_EXP")
	idExp, err := strconv.ParseInt(idTokenExp, 0, 64)
	if err != nil {
		fmt.Println("could not parse ID_TOKEN_EXP as int: %w", err)
	}

	authenHandler := handlers.AuthenHandler{
		TokenService: services.NewTokenService(&services.TSConfig{
			PrivKey:          privKey,
			PubKey:           pubKey,
			IDExpirationSecs: idExp,
		}),
	}
	authenv1 := engine.Group("/api/v1")
	{
		authen := authenv1.Group("/authen")
		{
			authen.POST("/LogIn", authenHandler.LogIn)
		}
	}

	engine.GET("/healthcheck", handlers.HealthCheck)
}
