package gocomposer

// isObject returns true if the JSON data starts with '{' and ends with '}'.
func isObject(data []byte) bool {
	if len(data) == 0 {
		return false
	}
	if data[0] != '{' || data[len(data)-1] != '}' {
		return false
	}
	return true
}

// isString returns true if the JSON data starts with '"' and ends with '"'.
func isString(data []byte) bool {
	if len(data) == 0 {
		return false
	}
	if data[0] != '"' || data[len(data)-1] != '"' {
		return false
	}
	return true
}

// isArray returns true if the JSON data starts with '[' and ends with ']'.
func isArray(data []byte) bool {
	if len(data) == 0 {
		return false
	}
	if data[0] != '[' || data[len(data)-1] != ']' {
		return false
	}
	return true
}
