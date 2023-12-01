package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func PostarProspects(respBody []map[string]interface{}) {
	crmBaseUrl := fmt.Sprintf("%spipefy", os.Getenv("API_URL"))

	jsonData, err := json.Marshal(respBody)
	if err != nil {
		fmt.Printf("Erro ao converter para JSON: %v\n", err)
		return
	}

	body := bytes.NewReader(jsonData)

	req, _ := http.NewRequest("POST", crmBaseUrl, body)
	req.Header = http.Header{
		"RoutineAuth":  {"G8xmGt0RU4e1GXkrr7EcDAUfhU9QGeyDzBqZybtyEK0PQ3u0btMwGtImwMBU4iDc"},
		"Content-Type": {"application/json"},
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Erro no post dos prospects: %v\n", err)
		return
	}
	defer resp.Body.Close()
}
