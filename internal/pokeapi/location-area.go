package pokeapi

type locationAreaResult struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type locationArea struct {
	Count    int                  `json:"count"`
	Next     *string              `json:"next"`
	Previous *string              `json:"previous"`
	Results  []locationAreaResult `json:"results"`
}
