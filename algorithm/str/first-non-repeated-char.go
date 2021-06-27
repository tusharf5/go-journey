package str

// Takes in a string and return first non-repeated character
func FirstNonRepeat(str string) string {
	lookup := make(map[string]int)
	var firstnonRepeatIndex int

	for index, value := range str {
		if _, ok := lookup[string(value)]; ok {
			if lookup[string(value)] == firstnonRepeatIndex {
				firstnonRepeatIndex = -1
			}
		} else {
			lookup[string(value)] = index
			if firstnonRepeatIndex == -1 {
				firstnonRepeatIndex = index
			}
		}
	}

	if firstnonRepeatIndex == -1 {
		return "❌"
	}

	return string(str[firstnonRepeatIndex])

}
