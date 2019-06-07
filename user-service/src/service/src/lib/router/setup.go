package router

import (
	"github.com/gin-gonic/gin"
	"os"
	"service/src/modules"
)

func Setup(r *gin.Engine) {
	rg := r.Group(os.Getenv("PATH_PREFIX"))

	modules.Group(rg)
}
