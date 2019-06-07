package schema

import "service/src/models/conn_log"

func SetupSchemas() {
	conn_log.Create()
}
