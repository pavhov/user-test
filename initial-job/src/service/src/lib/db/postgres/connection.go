package postgres

import (
	"github.com/go-pg/pg"
	"log"
)

func Connection() *pg.DB {
	s, err := Init()

	if err != nil {
		log.Fatal(err)
	}

	return s
}
