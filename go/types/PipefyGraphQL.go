package types

type Fields struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Node struct {
	ID           string   `json:"id"`
	Fields       []Fields `json:"fields"`
	CurrentPhase struct {
		Name string `json:"name"`
	} `json:"current_phase"`
}

type GraphQLResponse struct {
	Data struct {
		AllCards struct {
			Edges []struct {
				Node Node `json:"node"`
			} `json:"edges"`
		} `json:"allCards"`
	} `json:"data"`
}
