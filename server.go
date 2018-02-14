package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"

	"ichabod/controllers"
	"ichabod/db"

	"github.com/gin-gonic/gin"
)

//DefaultWriter ...
var DefaultWriter io.Writer = os.Stdout

//CORSMiddleware ...
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func main() {
	r := gin.Default()

	r.Use(gin.Logger())

	// store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	// r.Use(sessions.Sessions("ichabod-session", store))

	r.Use(CORSMiddleware())

	db.Init()

	v1 := r.Group("/v1")
	{
		application := new(controllers.ApplicationController)
		environment := new(controllers.EnvironmentController)

		v1.POST("/application", application.Create)
		v1.GET("/application/:appId", application.One)

		v1.POST("/application/:appId/environment", environment.Create)
		v1.GET("/application/:appId/environment/:slug", environment.One)
		v1.POST("/application/:appId/environment/:slug", environment.Update)

	}

	r.LoadHTMLGlob("./public/html/*")

	r.Static("/public", "./public")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"ichabodVersion": "v0.01",
			"goVersion":      runtime.Version(),
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})

	r.Run("localhost:3000")
}
