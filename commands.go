package main

type cliCommand struct {
	name        string
	description string
	callback    func(argument string) (string, error)
}

func commands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world. Each subsequent call to map should display the next 20 locations, and so on. The idea is that the map command will let us explore the world of Pokemon.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Similar to the map command, however, instead of displaying the next 20 locations, it displays the previous 20 locations. It's a way to go back.",
			callback:    commandMapb,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore",
			description: "see a list of all the Pokémon in a given area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "takes the name of a Pokemon as an argument. Catching Pokemon adds them to the Pokedex.",
			callback:    commandCatch,
		},
	}
}
