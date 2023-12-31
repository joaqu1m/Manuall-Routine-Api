package external

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Ociocos() *http.Response {
	crmBaseUrl := fmt.Sprintf("%scrm", os.Getenv("API_URL"))

	url := crmBaseUrl + "/ociosos"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header = http.Header{
		"RoutineAuth": {os.Getenv("ROUTINE_AUTH")},
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Erro na iteração pelo CRM: %v\n", err)
		return nil
	}

	return resp
}

func Heavy() *http.Response {
	crmBaseUrl := fmt.Sprintf("%scrm", os.Getenv("API_URL"))

	url := crmBaseUrl + "/heavy"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header = http.Header{
		"RoutineAuth": {os.Getenv("ROUTINE_AUTH")},
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Erro na iteração pelo CRM: %v\n", err)
		return nil
	}

	return resp
}

func Recentes() *http.Response {
	crmBaseUrl := fmt.Sprintf("%scrm", os.Getenv("API_URL"))

	url := crmBaseUrl + "/recentes"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header = http.Header{
		"RoutineAuth": {os.Getenv("ROUTINE_AUTH")},
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Erro na iteração pelo CRM: %v\n", err)
		return nil
	}

	return resp
}

func IniciarCrm(crmType string, id int) {
	crmBaseUrl := fmt.Sprintf("%scrm", os.Getenv("API_URL"))

	url := crmBaseUrl + "/" + crmType + "/iniciarCrm/" + strconv.Itoa(id)

	req, _ := http.NewRequest("POST", url, nil)
	req.Header = http.Header{
		"RoutineAuth": {os.Getenv("ROUTINE_AUTH")},
	}
	client := &http.Client{}
	_, err := client.Do(req)
	if err != nil {
		fmt.Printf("Erro na iteração pelo CRM: %v\n", err)
		return
	}
}
