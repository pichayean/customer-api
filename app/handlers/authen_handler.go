package handlers

import (
	"fmt"
	"net/http"

	"macus/models"
	"macus/services"

	"github.com/gin-gonic/gin"
)

type AuthenHandler struct {
	TokenService services.TokenService
}

func (h AuthenHandler) RegisterApi(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")
	{
		authen := v1.Group("/authen")
		{
			authen.POST("/LogIn", h.LogIn)
		}
	}
}

func (h AuthenHandler) LogIn(c *gin.Context) {
	var login models.LogIn
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		c.Abort()
		return
	}
	fmt.Println("HEre")
	accessToken, err := h.TokenService.NewToken(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		c.Abort()
		return
	}
	fmt.Println("HEre")
	c.JSON(http.StatusOK, accessToken)

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
	// 	Id:        login.UserName,
	// 	ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
	// })

	// if tokenSIgn, err := token.SignedString([]byte("MySignatureYSjoSWAQSF")); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": err.Error(),
	// 	})
	// } else {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"token": tokenSIgn,
	// 	})
	// }
}
