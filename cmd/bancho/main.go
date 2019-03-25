package main

import (
	"log"

	"git.iptq.io/nso/bancho"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	var err error

	config, err := bancho.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	server, err := bancho.NewInstance(&config)
	if err != nil {
		log.Fatal(err)
	}
	server.Run()
}
