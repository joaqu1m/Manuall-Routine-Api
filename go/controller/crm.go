package controller

import (
	"fmt"
	"manuall/routine-api/consumable"
	"manuall/routine-api/external"
)

func Crm() {
	fmt.Println("crm controller todo")

	consumable.Crm()
	external.Crm()
}
