package main

import (
	"flag"
	"log"

	"github.com/aidosgal/ei-jobs-core/cmd/api"
	"github.com/aidosgal/ei-jobs-core/config"
	"github.com/aidosgal/ei-jobs-core/database"
	"github.com/go-sql-driver/mysql"
)

func main() {
	shouldSeed := flag.Bool("seed", false, "Seed the database with initial data")
	flag.Parse()

	db, err := database.NewMySQLStorage(mysql.Config{
		User:   config.Envs.DBUser,
		Passwd: config.Envs.DBPassword,
		Addr:   config.Envs.DBAddress,
		DBName: config.Envs.DBName,

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

	if *shouldSeed {
		err = database.SeedDatabase(db)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Database seeded successfully")
	}

	server := api.NewAPIServer(config.Envs.Port, db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
