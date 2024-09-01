package main

// CompileDaemon -command="./simplecrudgo"
// postgresql://postgres:LTNvjuIsFHtuiQHjLHWFHYMoGicRalBM@autorack.proxy.rlwy.net:54592/railway
import (
	"simplecrudgo/controllers"
	"simplecrudgo/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostCreate)
	r.Run()
}
