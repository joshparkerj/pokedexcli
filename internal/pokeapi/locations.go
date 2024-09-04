package pokeapi

var currentPage int = 0

func locationsHelper() (result []string, err error) {
	la, err := getLocationArea(currentPage)
	if err != nil {
		return
	}

	for _, lar := range la.Results {
		result = append(result, lar.Name)
	}

	return
}

func Locations() (result []string, err error) {
	currentPage++
	result, err = locationsHelper()
	return
}

func LocationsBack() (result []string, err error) {
	currentPage--
	result, err = locationsHelper()
	return
}
