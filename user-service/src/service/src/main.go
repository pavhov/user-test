package main

import (
	"service/src/lib/config/env"
	"service/src/lib/db"
	"service/src/lib/router"
)

func main() {
	env.Init()
	db.Setup()

	r := router.Init()

	router.Setup(r)
	router.Run(r)
}
