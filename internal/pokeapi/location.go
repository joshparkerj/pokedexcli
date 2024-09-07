package pokeapi

type LocationGameIndex struct {
	GameIndex  int    `json:"game_index"`
	Generation Result `json:"generation"`
}

type LocationName struct {
	Language Result `json:"language"`
	Name     string `json:"name"`
}

type Location struct {
	Id          int                 `json:"id"`
	Name        string              `json:"name"`
	Areas       []Result            `json:"areas"`
	Region      Result              `json:"region"`
	Names       []LocationName      `json:"names"`
	GameIndices []LocationGameIndex `json:"game_indices"`
}
