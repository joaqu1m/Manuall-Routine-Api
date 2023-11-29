package external

import (
	"fmt"
	"net/http"
	"os"
)

var (
	pipefyBaseUrl = os.Getenv("API_URL") + "pipefy"
)

func Pipefy() {
	fmt.Println("pipefy external todo")

	resp, err := http.Get(pipefyBaseUrl)
	fmt.Println(resp.StatusCode)
	fmt.Println(err)
}
