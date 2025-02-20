package infrastructure

import (
	"os"

	"github.com/gin-gonic/gin"
)

func PrepareServer() *gin.Engine {
	SetupDatabase()
	return gin.Default()
}

func InitServer(engine *gin.Engine) {
	port := os.Getenv("PORT")
	engine.Run(":" + port)
}
