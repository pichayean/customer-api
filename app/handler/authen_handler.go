package handler

import (
	"net/http"
	"time"

	"macus/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type AuthenHandler struct {
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

// @Router /api/v1/customers [post]
func (h AuthenHandler) LogIn(c *gin.Context) {
	var login model.LogIn
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		c.Abort()
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		Id:        login.UserName,
		ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
	})

	if tokenSIgn, err := token.SignedString([]byte("MySignatureYSjoSWAQSF")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"token": tokenSIgn,
		})
	}
}
