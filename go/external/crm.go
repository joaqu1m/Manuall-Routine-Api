package external

import (
	"fmt"
	"net/http"
	"strconv"
)

const (
	baseUrl = "http://localhost:8080/routine/crm"
)

func Ociocos() *http.Response {
	url := baseUrl + "/ociosos"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header = http.Header{
		"RoutineAuth": {"G8xmGt0RU4e1GXkrr7EcDAUfhU9QGeyDzBqZybtyEK0PQ3u0btMwGtImwMBU4iDc"},
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
	url := baseUrl + "/heavy"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header = http.Header{
		"RoutineAuth": {"G8xmGt0RU4e1GXkrr7EcDAUfhU9QGeyDzBqZybtyEK0PQ3u0btMwGtImwMBU4iDc"},
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
	url := baseUrl + "/recentes"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header = http.Header{
		"RoutineAuth": {"G8xmGt0RU4e1GXkrr7EcDAUfhU9QGeyDzBqZybtyEK0PQ3u0btMwGtImwMBU4iDc"},
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

	url := baseUrl + "/" + crmType + "/iniciarCrm/" + strconv.Itoa(id)

	req, _ := http.NewRequest("POST", url, nil)
	req.Header = http.Header{
		"RoutineAuth": {"G8xmGt0RU4e1GXkrr7EcDAUfhU9QGeyDzBqZybtyEK0PQ3u0btMwGtImwMBU4iDc"},
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Erro na iteração pelo CRM: %v\n", err)
		return
	}

	fmt.Println(resp.StatusCode)
}
