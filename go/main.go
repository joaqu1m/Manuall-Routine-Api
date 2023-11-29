package main

import (
	"log"
	"manuall/routine-api/route"
	"net/http"
)

func main() {
	SetEnvs()
	route.Crm()
	route.Pipefy()

	log.Fatal(http.ListenAndServe(":3001", nil))
}
