package rpn

// InsertValues finds tokens that represent some column in the query
// and changes the token's value to corresponding data that came from the "map[csv-column]csv-value".
func InsertValues(values map[string]string, rpn []string) []string {

	for i, token := range rpn {
		if isVariable(token) {
			rpn[i] = values[token]
		}
	}
	return nil
}

// check if the token is a variable (column-name)
func isVariable(token string) bool {
	return false
}
