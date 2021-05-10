package parse

import (
	"errors"

	"github.com/seggga/csvquery/token"
)

var (
	errSelect error = errors.New("no columns has been chosen (section SELECT is empty)")
	errFrom   error = errors.New("no file has been chosen (section FROM is empty)")
)

type LexMachine struct {
	State   int
	Columns int // defines the number of columns that were mentioned it the query (SELECT and WHERE statements)
	Select  []token.Lexema
	From    []token.Lexema
	Where   []token.Lexema
}

// NewLexMachine fills up fields of the struct LexMachine
func NewLexMachine(queryLex []token.Lexema) (*LexMachine, error) {

	var lm LexMachine

	for _, lex := range queryLex {

		switch lex.Litera {
		case "SELECT":
			lm.State = 1
			continue
		case "FROM":
			lm.State = 2
			continue
		case "WHERE":
			lm.State = 3
			continue
		}

		switch lm.State {
		case 1:
			if lex.Token == "IDENT" {
				lm.Select = append(lm.Select, lex)
				lm.Columns += 1
			}
		case 2:
			if lex.Token == "STRING" {
				lm.From = append(lm.From, lex)
			}
		case 3:
			lm.Where = append(lm.Where, lex)
			if lex.Token == "IDENT" {
				lm.Columns += 1
			}
		}
	}

	// check if the query contains at least one column to be written to output
	if len(lm.Select) == 0 {
		return nil, errSelect
	}

	// check if the query contains at least one file to be read
	if len(lm.From) == 0 {
		return nil, errFrom
	}

	return &lm, nil

}
