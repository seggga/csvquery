package rpn

import "github.com/seggga/csvquery/token"

// InsertValues finds tokens that represent some column in the query
// and changes the token's value to corresponding data that came from the "map[csv-column]csv-value".
func InsertValues(values map[string]string, rpn []token.Lexema) []token.Lexema {

	// create a temp slice to leave the original rpn unchanged
	tempSlice := make([]token.Lexema, len(rpn))
	copy(tempSlice, rpn)

	for i, lex := range tempSlice {
		if isVariable(lex) {
			tempSlice[i].Litera = values[lex.Litera]
		}
	}
	return tempSlice
}

// check if the token is a variable (column-name)
func isVariable(lex token.Lexema) bool {

	return lex.Token == "IDENT"
}
