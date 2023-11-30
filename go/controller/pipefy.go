package controller

import (
	"fmt"
	"manuall/routine-api/consumable"
)

func Pipefy() {
	fmt.Println("pipefy controller todo")

	graphqlResponse := consumable.Pipefy()

	for _, edge := range graphqlResponse.Data.AllCards.Edges {
		fmt.Printf("Node ID: %s\n", edge.Node.ID)
		fmt.Printf("Node CurrentPhase: %s\n", edge.Node.CurrentPhase.Name)

		for _, field := range edge.Node.Fields {
			fmt.Println(fmt.Sprintf("FieldName: %s, FieldValue: %s", field.Name, field.Value))
		}
	}
}
