package conn_log

import (
	"github.com/gin-gonic/gin"
	"service/src/lib/sync"
)

func Router(r *gin.RouterGroup) {
	rg := r.Group("")

	rg.GET("/:first_user/:last_user", sync.DoOnceHandler(GetDupes))
}
