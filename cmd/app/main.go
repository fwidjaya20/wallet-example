package main

import (
	"fmt"
	"github.com/fwidjaya20/wallet-example/config"
	"github.com/fwidjaya20/wallet-example/internal/globals"
	"github.com/fwidjaya20/wallet-example/lib/database"
	"github.com/go-kit/kit/log"
	"github.com/jmoiron/sqlx"
	"github.com/oklog/oklog/pkg/group"
	"os"
)

func main() {
	var logger log.Logger
	var g group.Group

	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestamp)
	logger = log.With(logger, "caller", log.DefaultCaller)

	con := globals.DB()
	defer con.Close()

	initMigration(con)

	_ = logger.Log("exit", g.Run())
}

func initMigration(dbConn *sqlx.DB) {
	root, err := os.Getwd()
	if nil != err {
		panic(fmt.Sprintf("failed retrieve root path : %v", err.Error()))
	}

	migrationPath := fmt.Sprintf("%s/%s", root, config.GetEnv(config.MIGRATION_PATH))
	database.Migrate(dbConn.DB, config.GetEnv(config.DB_NAME), migrationPath)
}
