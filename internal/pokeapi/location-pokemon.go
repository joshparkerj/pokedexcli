package pokeapi

type ConditionValue struct {
	// TODO: what does the condition value look like?
}

type EncounterDetail struct {
	Chance          int              `json:"chance"`
	MaxLevel        int              `json:"max_level"`
	MinLevel        int              `json:"min_level"`
	Method          Result           `json:"method"`
	ConditionValues []ConditionValue `json:"condition_values"`
}

type PokemonEncounterVersionDetail struct {
	EncounterDetails []EncounterDetail `json:"encounter_details"`
	MaxChance        int               `json:"max_chance"`
	Version          Result            `json:"version"`
}

type PokemonEncounter struct {
	Pokemon        Result `json:"pokemon"`
	VersionDetails []PokemonEncounterVersionDetail
}

type EncounterMethodRateVersionDetail struct {
	MaxChance        int               `json:"max_chance"`
	Version          Result            `json:"version"`
	EncounterDetails []EncounterDetail `json:"encounter_details"`
}

type EncounterMethodRate struct {
	EncounterMethod Result                             `json:"encounter_method"`
	VersionDetails  []EncounterMethodRateVersionDetail `json:"version_details"`
}

type LocationPokemon struct {
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	GameIndex            int                   `json:"game_index"`
	Id                   int                   `json:"id"`
	Location             Result                `json:"location"`
	Name                 string                `json:"name"`
	Names                []LocationName        `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}
