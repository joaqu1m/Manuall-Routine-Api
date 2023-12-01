package consumable

import (
	"bytes"
	"encoding/json"
	"fmt"
	"manuall/routine-api/types"
	"net/http"
	"os"
	"time"
)

func Pipefy(pipeId string) types.GraphQLResponse {
	jsonMapInstance := map[string]string{
		"query": `
		{
			allCards(pipeId: ` + os.Getenv(pipeId) + `) {
				edges {
					node {
						id
						fields {
							name
							value
						},
						current_phase {
							name
						}
					}
				}
			}
		}
	`,
	}

	jsonResult, err := json.Marshal(jsonMapInstance)

	if err != nil {
		fmt.Printf("There was an error marshaling the JSON instance %v", err)
	}

	newRequest, err := http.NewRequest("POST", "https://api.pipefy.com/graphql", bytes.NewBuffer(jsonResult))
	newRequest.Header.Set("Content-Type", "application/json")
	newRequest.Header.Set("Authorization", "Bearer "+os.Getenv("PIPEFY_API_KEY"))

	client := &http.Client{Timeout: time.Second * 5}
	response, err := client.Do(newRequest)

	if err != nil {
		fmt.Printf("There was an error executing the request%v", err)
	}

	var graphqlResponse types.GraphQLResponse
	err = json.NewDecoder(response.Body).Decode(&graphqlResponse)
	if err != nil {
		fmt.Printf("Error decoding JSON response: %v", err)
	}

	return graphqlResponse
}
