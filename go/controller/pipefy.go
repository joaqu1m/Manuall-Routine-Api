package controller

import (
	"fmt"
	"manuall/routine-api/consumable"
	"manuall/routine-api/external"
	"manuall/routine-api/mockData"
	"strconv"
)

func Pipefy() {
	executarTipo(1)
	executarTipo(2)
}

func executarTipo(tipo int) {

	graphqlResponse := consumable.Pipefy()

	var prospects []map[string]interface{}

	for _, edge := range graphqlResponse.Data.AllCards.Edges {
		resultado, err := strconv.Atoi(edge.Node.ID)
		if err != nil {
			fmt.Printf("Erro ao converter ID: %v\n", err)
			return
		}

		prospectAtual := map[string]interface{}{
			"idCliente":   resultado,
			"status":      mockData.SyncField("status").Opts[edge.Node.CurrentPhase.Name].(int),
			"tipoUsuario": tipo,
		}

		for _, field := range edge.Node.Fields {
			fieldConfig := mockData.SyncField(field.Name)
			var valor interface{} = nil

			if fieldConfig.Opts != nil {
				valor = fieldConfig.Opts[field.Value]
			} else {
				valor = field.Value
			}

			if fieldConfig.Format != nil {
				valor = fieldConfig.Format(valor.(string))
			}

			prospectAtual[fieldConfig.DbCol] = valor
		}

		prospects = append(prospects, prospectAtual)
	}

	external.PostarProspects(prospects)
}
