package pokeapi

type Ability struct {
	IsHidden bool   `json:"is_hidden"`
	Slot     int    `json:"slot"`
	Ability  Result `json:"ability"`
}

type Cries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

type PokemonGameIndex struct {
	GameIndex int    `json:"game_index"`
	Version   Result `json:"version"`
}

type VersionGroupDetail struct {
	LevelLearnedAt  int    `json:"level_learned_at"`
	MoveLearnMethod Result `json:"move_learn_method"`
	VersionGroup    Result `json:"version_group"`
}

type Move struct {
	Move                Result               `json:"move"`
	VersionGroupDetails []VersionGroupDetail `json:"version_group_details"`
}

type Stat struct {
	BaseStat int    `json:"base_stat"`
	Effort   int    `json:"effort"`
	Stat     Result `json:"stat"`
}

type Type struct {
	Slot int    `json:"slot"`
	Type Result `json:"type"`
}

type Pokemon struct {
	BaseExperience         int                `json:"base_experience"`
	Height                 int                `json:"height"`
	HeldItems              []struct{}         `json:"held_items"`
	Id                     int                `json:"id"`
	IsDefault              bool               `json:"is_default"`
	LocationAreaEncounters string             `json:"location_area_encounters"`
	Name                   string             `json:"name"`
	Order                  int                `json:"order"`
	PastAbilities          []struct{}         `json:"past_abilities"`
	PastTypes              []struct{}         `json:"past_types"`
	Weight                 int                `json:"weight"`
	Forms                  []Result           `json:"forms"`
	Species                Result             `json:"species"`
	Abilities              []Ability          `json:"abilities"`
	Cries                  Cries              `json:"cries"`
	GameIndices            []PokemonGameIndex `json:"game_indices"`
	Moves                  []Move             `json:"moves"`
	Sprites                Sprites            `json:"sprites"`
	Stats                  []Stat             `json:"stats"`
	Types                  []Type             `json:"types"`
}
