package main

// CompileDaemon -command="./simplecrudgo"
// postgresql://postgres:LTNvjuIsFHtuiQHjLHWFHYMoGicRalBM@autorack.proxy.rlwy.net:54592/railway
import (
	"simplecrudgo/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
