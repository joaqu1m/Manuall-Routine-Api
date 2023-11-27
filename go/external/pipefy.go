package external

import (
	"fmt"
	"net/http"
)

func Pipefy() {
	fmt.Println("pipefy external todo")

	url := "http://localhost:8080/routine/pipefy"

	resp, err := http.Get(url)
	fmt.Println(resp.StatusCode)
	fmt.Println(err)
}
