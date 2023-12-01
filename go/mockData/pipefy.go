package mockData

import (
	"manuall/routine-api/types"
	"regexp"
)

var (
	fields = map[string]types.Field{
		"status": {
			DbCol: "status",
			Opts: map[string]interface{}{
				"Lead":              1,
				"Oportunidade":      2,
				"Cliente Adquirido": 3,
				"Cliente Perdido":   4,
			},
			Format: nil,
		},
		"A partir de qual canal você chegou aqui?": {
			DbCol: "optCanal",
			Opts: map[string]interface{}{
				"Redes Sociais": 1,
				"Pesquisa":      2,
				"Indicação":     3,
				"Próprio site":  4,
			},
			Format: nil,
		},
		"Nome Completo": {
			DbCol:  "nome",
			Opts:   nil,
			Format: nil,
		},
		"Email": {
			DbCol:  "email",
			Opts:   nil,
			Format: nil,
		},
		"Telefone": {
			DbCol: "fone",
			Opts:  nil,
			Format: func(value string) string {
				re := regexp.MustCompile("[^0-9]+")
				formattedValue := re.ReplaceAllString(value[4:], "")

				return formattedValue
			},
		},
		"Você mora em:": {
			DbCol: "optCidade",
			Opts: map[string]interface{}{
				"São Paulo":             1,
				"São Bernardo do Campo": 2,
				"São Caetano do Sul":    3,
				"Santo André":           4,
				"Osasco":                5,
				"Bauru":                 6,
				"Outro":                 7,
			},
			Format: nil,
		},
		"Já conhece a Manuall?": {
			DbCol: "blnConheceManuall",
			Opts: map[string]interface{}{
				"Sim": true,
				"Não": false,
			},
			Format: nil,
		},
		"Qual(is) dessas áreas de serviços você está buscando?": {
			DbCol: "areaId",
			Opts: map[string]interface{}{
				"Jardineiro":  1,
				"Pintor":      2,
				"Eletricista": 3,
				"Encanador":   4,
				"Marceneiro":  5,
				"Montador":    6,
				"Gesseiro":    7,
				"Nenhuma":     nil,
			},
			Format: nil,
		},
		"Você teria interesse em aprender algum dos serviços citados anteriormente?": {
			DbCol: "blnAprender",
			Opts: map[string]interface{}{
				"Sim, possuo interesse.":                              true,
				"Não, quero apenas contratar o prestador de serviço.": false,
			},
			Format: nil,
		},
		"Você tem interesse pela Manuall?": {
			DbCol: "blnInteresseManuall",
			Opts: map[string]interface{}{
				"Sim": true,
				"Não": false,
			},
			Format: nil,
		},
		"Selecione a sua área de serviço de interesse:": {
			DbCol: "areaId",
			Opts: map[string]interface{}{
				"Jardineiro":  1,
				"Pintor":      2,
				"Eletricista": 3,
				"Encanador":   4,
				"Marceneiro":  5,
				"Montador":    6,
				"Gesseiro":    7,
				"Nenhuma":     nil,
			},
			Format: nil,
		},
		"Você teria interesse em ensinar um pouco sobre a sua área ao outro?": {
			DbCol: "blnInteresseEnsinar",
			Opts: map[string]interface{}{
				"Sim, possuo interesse.":                 true,
				"Não, quero apenas prestar meu serviço.": false,
			},
			Format: nil,
		},
		"Utilizou o cupom e se tornou um cliente Contratante da Manuall?": {
			DbCol: "blnCupom",
			Opts: map[string]interface{}{
				"Sim": true,
				"Não": false,
			},
			Format: nil,
		},
		"Utilizou o cupom e se tornou um cliente Prestador de Serviço da Manuall?": {
			DbCol: "blnCupom",
			Opts: map[string]interface{}{
				"Sim": true,
				"Não": false,
			},
			Format: nil,
		},
		"Explique sua falta de interesse pela Manuall": {
			DbCol:  "msgDesistencia",
			Opts:   nil,
			Format: nil,
		},
	}
)

func SyncField(fieldName string) types.Field {
	return fields[fieldName]
}
