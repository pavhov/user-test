package conn_log

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"log"
	"service/src/lib/db/postgres"
	"time"
)

var ConnLogSchema struct {
	Schema *pg.DB
	//Fields *bson.M
	//Indexes *bson.M
	//Types  struct {
	//	Output OutputSchema
	//	Create CreateSchema
	//	Update UpdateSchema
	//}
}


type ConnLog struct {
	UserId int64 `sql:"alias:user_id,type:bigint"`
	IpAddr string `sql:"alias:ip_addr,type:varchar(15)"`
	Ts time.Time `sql:"alias:ts,type:timestamp"`
}


/** Browser model */
var connection *pg.DB

/** Model builder */
func Create() {
	if connection == nil {
		connection := *postgres.Connection()
		if err := connection.CreateTable((*ConnLog)(nil), &orm.CreateTableOptions{
			IfNotExists: true,
		}); err != nil {
			log.Fatal(err)
		}
		if _, err := connection.Exec("create index if not exists conn_logs_ip_addr_index on conn_logs (ip_addr); " +
			"create index if not exists conn_logs_user_id_index on conn_logs (user_id); "); err != nil {
			log.Fatal(err)
		}
		ConnLogSchema.Schema = &connection
	}
}
