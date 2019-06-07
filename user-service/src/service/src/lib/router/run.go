package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func Run(r *gin.Engine) {
	if err := r.Run(":" + os.Getenv("HTTP_PORT")); err != nil {
		log.Fatal(err)
	}
}
