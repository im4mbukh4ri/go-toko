package controllers

import (
	"fmt"
	"log"
	"os"

	"github.com/im4mbukh4ri/go-toko/app/models"
	"github.com/im4mbukh4ri/go-toko/database/seeders"
	"github.com/urfave/cli"
)

func (server *Server) dbMigrate() {
	for _, model := range models.RegisterModels() {
		err := server.DB.Debug().AutoMigrate(model.Model)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Database Migrated Success")
}

func (server *Server) InitCommands(config AppConfig, dbConfig DBConfig) {
	server.InitializeDB(dbConfig)
	cmdApp := cli.NewApp()
	cmdApp.Commands = []cli.Command{
		{
			Name: "db:migrate",
			Action: func(c *cli.Context) error {
				server.dbMigrate()
				return nil
			},
		},
		{
			Name: "db:seed",
			Action: func(c *cli.Context) error {
				err := seeders.DBSeed(server.DB)
				if err != nil {
					log.Fatal(err)
				}

				return nil
			},
		},
	}
	err := cmdApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
