package api

import (
	"net/http"
	"os"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"

	"github.com/ravaj-group/toolbelt/daemon/api/request"
)

func Listen() {
	server := martini.Classic()

	server.Use(jsonRequest)
	registerRoutes(server)

	server.RunOnAddr(":" + os.Getenv("TOOLBELT_API_PORT"))
}

func registerRoutes(server *martini.ClassicMartini) {
	server.Post("/exec", binding.Bind(request.CmdRequest{}), Execute)
}

func jsonRequest(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Content-Type") != "application/json" {
		res.WriteHeader(http.StatusBadRequest)
		res.Header().Set("Content-Type", "application/json")
		res.Write([]byte("{\"error\":\"Content-Type specified must be application/json\"}"))
	}
}
