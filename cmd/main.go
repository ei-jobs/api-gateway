package main

import (
	"log"

	"github.com/aidosgal/ei-jobs-core/cmd/api"
	"github.com/aidosgal/ei-jobs-core/config"
	"github.com/aidosgal/ei-jobs-core/database"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := database.NewMySQLStorage(mysql.Config{
		User:   config.Envs.DBUser,
		Passwd: config.Envs.DBPassword,
		Addr:   config.Envs.DBAddress,

		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	err = database.Migrate(db)
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(config.Envs.Port, db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
