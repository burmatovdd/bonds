package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	helper "modules/helpers"
	"net/http"
)

type MiddlewareAuth struct {
	Res bool   `json:"res"`
	Id  string `json:"id"`
}

// Authz validates token and authorizes users
func Authentication(token string) MiddlewareAuth {
	res := true
	var middlewareAuth MiddlewareAuth
	if token == "" {
		fmt.Println("no token!")
		res = false
	}

	claims, err := helper.ValidateToken(token)
	if err != "" {
		res = false
		fmt.Println("err: ", err)
	}

	middlewareAuth.Res = res
	middlewareAuth.Id = claims.Uid
	fmt.Println("middlewareAuth:", middlewareAuth)
	return middlewareAuth
}
func Middleware1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("i'm in Middleware1")
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization header provided")})
			c.Abort()
			return
		}

		claims, err := helper.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		if claims.Uid == "" {
			c.AbortWithStatus(401)
			return
		}

		c.Set("uid", claims.Uid)

		c.Next()

	}
}
