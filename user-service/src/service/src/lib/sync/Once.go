package sync

import (
	"github.com/gin-gonic/gin"
	"sync"
)

func DoOnceHandler(fn func(c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var once sync.Once
		done := make(chan bool)
		go func() {
			once.Do(func() {
				fn(c)
			})
			done <- true
		}()
		<-done
	}
}
