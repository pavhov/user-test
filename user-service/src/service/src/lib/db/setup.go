package db

import "service/src/lib/db/postgres/schema"

func Setup() {
	schema.SetupSchemas()
}
