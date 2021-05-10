package parse

import (
	"errors"

	"github.com/seggga/csvquery/token"
)

// CheckCols checks that columns in SELECT statement present in the given slice
func CheckCols(cols []string, lm *LexMachine) error {

	// check SELECT statement
	counter := len(lm.Select)
	for _, colInQuery := range lm.Select {
		for _, colInTable := range cols {
			if colInQuery.Litera == colInTable {
				counter--
				break
			}
		}
	}
	if counter > 0 {
		return errors.New("wrong SELECT statement: some columns were not found in the file")
	}

	// extract colunms frome WHERE statement
	var colsInWhere []token.Lexema
	for _, colInQuery := range lm.Where {
		if colInQuery.Token == "IDENT" {
			colsInWhere = append(colsInWhere, colInQuery)
		}
	}

	// check WHERE statement
	counter = len(colsInWhere)
	for _, colInQuery := range colsInWhere {
		for _, colInTable := range cols {
			if colInQuery.Litera == colInTable {
				counter--
				break
			}
		}
	}

	if counter > 0 {
		return errors.New("wrong WHERE statement: some columns were not found in the file")
	}

	return nil
}
