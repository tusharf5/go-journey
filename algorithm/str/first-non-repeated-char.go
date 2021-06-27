package str

// Takes in a string and return first non-repeated character
func FirstNonRepeat(str string) string {
	multiLookup := make(map[string]int)
	lookup := make(map[string]int)

	for index, value := range str {

		if _, ok := multiLookup[string(value)]; ok {

			continue

		} else if _, ok := lookup[string(value)]; ok {

			multiLookup[string(value)] = lookup[string(value)]
			delete(lookup, string(value))

		} else {
			lookup[string(value)] = index

		}
	}

	firstIndex := -1

	for _, index := range lookup {
		if firstIndex == -1 {
			firstIndex = index
		}

		if index < firstIndex {
			firstIndex = index
		}
	}

	if firstIndex == -1 {
		return "❌"
	}

	return string(str[firstIndex])

}
