package handlers

import (
	"fmt"
	"io/ioutil"
	"macus/database"
	"macus/middlewares"
	"macus/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func AddAPIs(db *database.GormStore, engine *gin.Engine) {
	customersHandler := NewCustomerHandler(db)
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

	// privKeyFile := os.Getenv("PRIV_KEY_FILE")
	privKeyFile := "./config/rsa_private.pem"
	priv, err := ioutil.ReadFile(privKeyFile)
	if err != nil {
		fmt.Println("could not read private key pem file: %w", err)
	}
	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(priv)
	fmt.Println(privKey)
	if err != nil {
		fmt.Println("could not parse private key: %w", err)
	}

	pubKeyFile := "./config/psa_public.pem"
	pub, err := ioutil.ReadFile(pubKeyFile)
	if err != nil {
		fmt.Println("could not read public key pem file: %w", err)
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pub)
	if err != nil {
		fmt.Println("could not parse public key: %w", err)
	}

	// load expiration lengths from env variables and parse as int
	// idTokenExp := os.Getenv("ID_TOKEN_EXP")
	idTokenExp := "900"
	idExp, err := strconv.ParseInt(idTokenExp, 0, 64)
	if err != nil {
		fmt.Println("could not parse ID_TOKEN_EXP as int: %w", err)
	}

	authenHandler := AuthenHandler{
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

	engine.GET("/healthcheck", HealthCheck)
}

func AddSwagger(engine *gin.Engine) {
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
