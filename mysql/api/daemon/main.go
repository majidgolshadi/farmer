package main

import (
	"github.com/ravaj-group/farmer/mysql/api/daemon/db"
	"github.com/ravaj-group/farmer/mysql/api/daemon/api"
)

func main() {
	db.Connect()
	api.Listen()

	defer db.Close()
}
