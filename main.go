package main

import (
	"golang-api/cmd"
	"golang-api/pkg/db"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var server = cmd.Server{}

func main() {

	server.Init()

	db.Init()

	db.DBLoad(db.DB)

	server.InitializeRoutes()

	log.Infoln("Listening on network " + server.NetworkAddress)
	log.Fatal(http.ListenAndServe(server.NetworkAddress, server.Router))

}
