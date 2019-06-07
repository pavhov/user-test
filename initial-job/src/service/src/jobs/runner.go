package jobs

import "service/src/jobs/conn_log"

func Setup() {
	conn_log.Run()
}
