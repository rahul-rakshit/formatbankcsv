package utils

func HasColumns(headerColumn, requiredColumns []string) bool {
	for _, required := range requiredColumns {
		found := false

		for _, header := range headerColumn {
			if header == required {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}
