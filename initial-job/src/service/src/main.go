package main

import (
	"service/src/jobs"
	"service/src/lib/config/env"
	"service/src/lib/db"
)

func main() {
	env.Init()
	db.Setup()
	jobs.Setup()
}
