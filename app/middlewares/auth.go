package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"

	// "github.com/go-playground/validator"
	"macus/model/apperrors"
)

// IMPORTS OMITTED - Make sure to import validator/v10
// My auto import always uses V9

type authHeader struct {
	IDToken string `header:"Authorization"`
}

// used to help extract validation errors
type invalidArgument struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

// AuthUser extracts a user from the Authorization header
// which is of the form "Bearer token"
// It sets the user to the context if the user exists
func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := authHeader{}

		// bind Authorization Header to h and check for validation errors
		if err := c.ShouldBindHeader(&h); err != nil {
			if errs, ok := err.(validator.ValidationErrors); ok {
				// we used this type in bind_data to extract desired fields from errs
				// you might consider extracting it
				var invalidArgs []invalidArgument

				for _, err := range errs {
					invalidArgs = append(invalidArgs, invalidArgument{
						err.Field(),
						err.Value().(string),
						err.Tag(),
						err.Param(),
					})
				}

				err := apperrors.NewBadRequest("Invalid request parameters. See invalidArgs")

				c.JSON(err.Status(), gin.H{
					"error":       err,
					"invalidArgs": invalidArgs,
				})
				c.Abort()
				return
			}

			// otherwise error type is unknown
			err := apperrors.NewInternal()
			c.JSON(err.Status(), gin.H{
				"error": err,
			})
			c.Abort()
			return
		}

		idTokenHeader := strings.Split(h.IDToken, "Bearer ")

		if len(idTokenHeader) < 2 {
			err := apperrors.NewAuthorization("Must provide Authorization header with format `Bearer {token}`")

			c.JSON(err.Status(), gin.H{
				"error": err,
			})
			c.Abort()
			return
		}

		// validate ID token here
		// user, err := s.ValidateIDToken(idTokenHeader[1])

		// if err != nil {
		// 	err := apperrors.NewAuthorization("Provided token is invalid")
		// 	c.JSON(err.Status(), gin.H{
		// 		"error": err,
		// 	})
		// 	c.Abort()
		// 	return
		// }

		// c.Set("user", user)

		c.Next()
	}
}

// https://github.com/JacobSNGoodwin/memrizr/tree/master/account/model/apperrors
