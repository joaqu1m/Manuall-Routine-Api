package consumable

import (
	"bytes"
	"encoding/json"
	"fmt"
	"manuall/routine-api/types"
	"net/http"
	"time"
)

func Pipefy() types.GraphQLResponse {
	jsonMapInstance := map[string]string{
		"query": `
        {
            allCards(pipeId: 303175914) {
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
	newRequest.Header.Set("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJ1c2VyIjp7ImlkIjozMDI0MzEwODcsImVtYWlsIjoibWFudWFsbC5zZXJ2aWNlc0BvdXRsb29rLmNvbSIsImFwcGxpY2F0aW9uIjozMDAyNDUyMDZ9fQ.GQElgbfKYBMeyz1F3kLOjbH-nUsMbD2BswiL6NBU9xqlXpFHq5FV645zzO6vEj1PRqV_8APJ7o5ig6hMICbwZw")

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
