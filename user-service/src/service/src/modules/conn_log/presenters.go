package conn_log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"service/src/models/conn_log"
	"time"
)

func GetDupes(c *gin.Context) {
	firstUser := c.Param("first_user")
	lastUser := c.Param("last_user")
	var res []conn_log.ConnLog

	where1 := conn_log.ConnLogSchema.Schema.Model(&conn_log.ConnLog{}).Where("user_id = ?", firstUser)
	where2 := where1.Where("ip_addr in (select ip_addr from conn_logs as c where c.user_id = ?)", lastUser)
	query := where2.Column("user_id", "ip_addr").Group("user_id", "ip_addr").Having("count(*) > 1")
	startTime := time.Now().Second()
	if err := query.Select(&res); err != nil {
		 c.JSON(400, gin.H{"status": false, "error": err.Error()})
		return
	}
	endTime := time.Now().Second()
	fmt.Println(endTime- startTime, "second")
	if res == nil {
		c.JSON(200, gin.H{"dupes": false})
		return
	}
	c.JSON(200, gin.H{"dupes": true})
}
