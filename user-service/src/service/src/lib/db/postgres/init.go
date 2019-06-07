package postgres

import (
	"fmt"
	"os"

	"github.com/go-pg/pg"
)

type dbLogger struct { }

func (d dbLogger) BeforeQuery(q *pg.QueryEvent) {
	//fmt.Println(q.FormattedQuery())
}

func (d dbLogger) AfterQuery(q *pg.QueryEvent) {
	fmt.Println(q.FormattedQuery())
}

func Init() (*pg.DB, error) {
	c := pg.Connect(&pg.Options{Addr: os.Getenv("POSTGRES"), Database: os.Getenv("POSTGRES_GB"), User: os.Getenv("POSTGRES_USER")})
	c.AddQueryHook(dbLogger{})
	_, err := c.Exec("SELECT 1")
	return c, err
}
