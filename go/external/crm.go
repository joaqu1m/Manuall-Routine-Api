package external

import (
	"fmt"
	"net/http"
)

func Crm() {
	fmt.Println("crm external todo")

	url := "http://localhost:8080/routine/crm"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Erro na iteração pelo CRM: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
}
