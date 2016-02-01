package main

import (
	"github.com/ravaj-group/farmer/toolbelt_daemon/daemon/api"
	"github.com/ravaj-group/farmer/toolbelt_daemon/daemon/db"
)


func main() {
	api.Listen()
	db.Connect()

	defer db.Close()
}
