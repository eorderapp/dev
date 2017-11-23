package main

import (
	"github.com/gin-gonic/gin"
	"github.com/appleboy/gin-jwt"
	"net/http"
	"time"
)

//http://blog.theodo.fr/2016/11/securize-a-koa-api-with-a-jwt-token/

func main() {
	router := gin.Default()
	auth := setAuthHandler(router)
	addHandlers(auth)
	http.ListenAndServe(":8080", router)
}




func setAuthHandler(router *gin.Engine) *gin.RouterGroup {
	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        []byte("secret key"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {
			if (userId == "admin" && password == "admin") || (userId == "test" && password == "test") {
				return userId, true
			}

			return userId, false
		},
		Authorizator: func(userId string, c *gin.Context) bool {
			if userId == "admin" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header:Authorization",
		TokenHeadName: "Eorder",
		TimeFunc:      time.Now,
	}

	router.POST("/login", authMiddleware.LoginHandler)

	auth := router.Group("/")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", helloHandler)
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	}
	return auth
}

func addHandlers(auth *gin.RouterGroup) {

	auth.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	auth.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

}

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"userID": claims["id"],
		"text":   "Hello World.",
	})
}
