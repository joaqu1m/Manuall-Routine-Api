package main

import (
	"fmt"
	"manuall/routine-api/consumable"
	"manuall/routine-api/controller"
	"manuall/routine-api/external"
	"manuall/routine-api/route"
	"net/http"
)

func main() {

	url := "http://localhost:8080/routine/crm"

	resp, err := http.Get(url)
	fmt.Println(resp.StatusCode)
	fmt.Println(err)

	external.Crm()
	consumable.Crm()
	controller.Crm()
	route.Crm()
}
