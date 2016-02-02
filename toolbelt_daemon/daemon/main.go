package main

import (
	"github.com/ravaj-group/farmer/toolbelt_daemon/daemon/api"
	"github.com/ravaj-group/farmer/toolbelt_daemon/daemon/db"
)

func main() {
	db.Connect()
	api.Listen()
}
