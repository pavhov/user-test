package conn_log

import (
	"fmt"
	"log"
	"math/rand"
	"service/src/models/conn_log"
	"time"
)

func Run() {
	ips := []string{
		"127.0.0.1",
		"127.0.0.2",
		"127.0.0.3",
		"127.0.0.4",
		"127.0.0.5",
		"127.0.0.6",
		"127.0.0.7",
		"127.0.0.8",
		"127.0.0.9",
		"127.0.0.10",
		"127.0.0.11",
		"127.0.0.12",
		"127.0.0.13",
		"127.0.0.14",
		"127.0.0.15",
		"127.0.0.16",
		"127.0.0.17",
		"127.0.0.18",
		"127.0.0.19",
		"127.0.0.20",
	}

	var userslen int
	baseiter := 20
	fmt.Println("Start import users")
	for i := 0; i < baseiter; i++ {
		usrlen := 1000000
		userslen = (i + 1) * usrlen
		users := make([]conn_log.ConnLog, usrlen)
		for l := 0; l < usrlen; l++ {
			id := rand.Intn(10)
			users[l] = conn_log.ConnLog{
				UserId: int64(i + 1),
				IpAddr: ips[id],
				Ts:     time.Now(),
			}
		}
		if err := conn_log.ConnLogSchema.Schema.Insert(&users); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Completed import users", userslen)
	}
	users := make([]conn_log.ConnLog, 2)
	users[0] = conn_log.ConnLog{
		UserId: int64(baseiter + 1),
		IpAddr: "127.0.0.21",
		Ts:     time.Now(),
	}
	users[1] = conn_log.ConnLog{
		UserId: int64(baseiter + 2),
		IpAddr: "127.0.0.22",
		Ts:     time.Now(),
	}
	if err := conn_log.ConnLogSchema.Schema.Insert(&users); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Task completed")
}
