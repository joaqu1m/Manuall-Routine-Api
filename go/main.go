package main

import (
	"log"
	"manuall/routine-api/route"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	SetEnvs()
	route.Crm(router)
	route.Pipefy(router)

	log.Fatal(http.ListenAndServe(":3001", router))
}
