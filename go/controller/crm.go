package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"manuall/routine-api/consumable"
	"manuall/routine-api/external"
	"net/http"
)

type Usuario struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

func Crm() {
	executarCrm(external.Ociocos(), "ociosos")
	executarCrm(external.Heavy(), "heavy")
	executarCrm(external.Recentes(), "recentes")
}

func executarCrm(resp *http.Response, tipo string) {
	if resp.StatusCode != 200 {
		return
	}

	ociososData, err := processResponse(resp)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error processing %s response:", tipo), err)
		return
	}
	for _, item := range ociososData {
		external.IniciarCrm(tipo, item.ID)
		consumable.Crm(item.Email)
	}
}

func processResponse(resp *http.Response) ([]Usuario, error) {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response body: %v", err)
	}

	var data []Usuario
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("Error decoding JSON: %v", err)
	}

	return data, nil
}
