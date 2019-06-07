package modules

import (
	"github.com/gin-gonic/gin"
	"service/src/modules/conn_log"
)

func Group(r *gin.RouterGroup) {
	conn_log.Router(r)
}
