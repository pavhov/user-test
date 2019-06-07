package router

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/szuecs/gin-glog"
	"time"
)

func Init() (g *gin.Engine) {
	flag.Parse()
	r := gin.New()

	r.Use(ginglog.Logger(1 * time.Millisecond))
	r.Use(gin.Recovery())

	return r
}
