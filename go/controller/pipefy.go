package controller

import (
	"fmt"
	"manuall/routine-api/consumable"
	"manuall/routine-api/external"
)

func Pipefy() {
	fmt.Println("pipefy controller todo")

	consumable.Pipefy()
	external.Pipefy()
}
