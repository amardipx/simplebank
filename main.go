package main

import (
	"database/sql"
	"log"

	"github.com/amardipx/simplebank/api"
	db "github.com/amardipx/simplebank/db/sqlc"
	"github.com/amardipx/simplebank/util"
	_ "github.com/lib/pq" // needed for postgres driver
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
