package parse

import (
	"fmt"
	"strings"
)

// PrintLine produces the output
func PrintLine(valuesMap map[string]string, lm *LexMachine) {

	for _, lex := range lm.Select {
		fmt.Printf("%s\t", valuesMap[lex.Litera])
	}
	fmt.Println()
}

// PrintHeader produces the output
func PrintHeader(lm *LexMachine) {

	for _, lex := range lm.Select {
		fmt.Printf("%s\t", lex.Litera)
	}
	fmt.Println()
}

// FillTheMap produces a map with data of the current row: map[column_name]column_value
func FillTheMap(header, data []string, lm *LexMachine) map[string]string {

	valuesMap := make(map[string]string, lm.Columns)

	// fill the output map with SELECT-data
	for _, lex := range lm.Select {
		for i := 0; i < len(header); i += 1 {
			if strings.EqualFold(lex.Litera, header[i]) {
				valuesMap[lex.Litera] = data[i]
			}
		}
	}

	// add WHERE-data to the output map
	for _, lex := range lm.Where { //подход не учитывает скобки. Надо исправлять
		if lex.Token != "IDENT" {
			continue
		}
		for i := 0; i < len(header); i += 1 {
			if strings.EqualFold(lex.Litera, header[i]) {
				valuesMap[lex.Litera] = data[i]
			}
		}
	}
	return valuesMap
}
