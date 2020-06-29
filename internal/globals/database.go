package globals

import (
	"context"
	"fmt"
	"github.com/fwidjaya20/wallet-example/config"
	"github.com/fwidjaya20/wallet-example/lib/database"
	"github.com/jmoiron/sqlx"
	"log"
	"sync"
)

var db *sqlx.DB
var once sync.Once

func DB() *sqlx.DB {
	once.Do(func() {
		var err error

		conn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			config.GetEnv(config.DB_HOST),
			config.GetEnv(config.DB_PORT),
			config.GetEnv(config.DB_USER),
			config.GetEnv(config.DB_PASS),
			config.GetEnv(config.DB_NAME),
		)

		db, err = sqlx.Connect(config.GetEnv(config.DB_DRIVER), conn)

		if nil != err {
			log.Fatal(err)
		}
	})
	return db
}

func GetDefaultLimit() int64 {
	return int64(10)
}

func GetQuery(ctx context.Context) *database.Queryable {
	q, ok := database.QueryFromContext(ctx)
	if !ok {
		panic("values when get query from context. please using transaction")
	}
	return &q
}
