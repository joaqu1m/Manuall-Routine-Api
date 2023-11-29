package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"manuall/routine-api/external"
	"net/http"
)

type Usuario struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

func Crm() {
	ociososResp := external.Ociocos()
	ociososData, err := processResponse(ociososResp)
	if err != nil {
		fmt.Println("Error processing Ociosos response:", err)
		return
	}
	for _, item := range ociososData {
		external.IniciarCrm("ociosos", item.ID)
	}

	heavyResp := external.Heavy()
	heavyData, err := processResponse(heavyResp)
	if err != nil {
		fmt.Println("Error processing Ociosos response:", err)
		return
	}
	for _, item := range heavyData {
		external.IniciarCrm("heavy", item.ID)
	}

	recentesResp := external.Recentes()
	recentesData, err := processResponse(recentesResp)
	if err != nil {
		fmt.Println("Error processing Ociosos response:", err)
		return
	}
	for _, item := range recentesData {
		external.IniciarCrm("recentes", item.ID)
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
