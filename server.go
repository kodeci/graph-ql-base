package main
import "github.com/gin-gonic/gin"
func main() {
    // Creates a router without any middleware by default
    r := gin.New()
    // By default gin.DefaultWriter = os.Stdout
    r.Use(gin.Logger())
    // Recovery middleware recovers from any panics and writes a 500 if there was one.
    r.Use(gin.Recovery())
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    // Listen and serve on 0.0.0.0:8080
    r.Run(":8080")
}