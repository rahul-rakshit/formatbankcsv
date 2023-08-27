package utils

func IsLineEqual(line1 []string, line2 []string) bool {
	if len(line1) != len(line2) {
		return false
	}

	for i, v := range line1 {
		if v != line2[i] {
			return false
		}
	}

	return true
}
