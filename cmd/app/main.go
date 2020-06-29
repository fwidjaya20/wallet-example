package main

import (
	"fmt"
	"github.com/fwidjaya20/wallet-example/cmd/container"
	"github.com/fwidjaya20/wallet-example/cmd/http"
	"github.com/fwidjaya20/wallet-example/config"
	"github.com/fwidjaya20/wallet-example/internal/globals"
	"github.com/fwidjaya20/wallet-example/lib/database"
	"github.com/go-chi/chi"
	"github.com/go-kit/kit/log"
	"github.com/jmoiron/sqlx"
	"github.com/oklog/oklog/pkg/group"
	"github.com/rs/cors"
	netHttp "net/http"
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

	di := container.New(logger)

	initMigration(con)
	initHTTP(logger, &g, di)

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

func initHTTP(
	logger log.Logger,
	g *group.Group,
	container container.Container) {
	_ = logger.Log(logger, "Component", "HTTP")

	HTTP_ADDR := config.GetEnv(config.HTTP_ADDR)

	if len(HTTP_ADDR) < 1 {
		panic(fmt.Sprintf("Environment Missing!\n*%s* is required", HTTP_ADDR))
	}

	var router *chi.Mux
	router = chi.NewRouter()
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	router.Use(corsHandler.Handler)
	router.Mount("/v1", http.MakeHandler(router, container))

	server := &netHttp.Server{
		Addr:    HTTP_ADDR,
		Handler: router,
	}

	g.Add(
		func() error {
			_ = logger.Log("transport", "debug/HTTP", "addr", HTTP_ADDR)
			return server.ListenAndServe()
		},
		func(err error) {
			if nil != err {
				_ = logger.Log("transport", "debug/HTTP", "addr", HTTP_ADDR, "values", err)
				panic(err)
			}
		},
	)
}
