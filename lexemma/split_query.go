package lexemma

import "github.com/seggga/csvquery/token"

func SplitQuery(query string) []Lexemma {

	var scanner token.Scanner
	var lexx []Lexemma

	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(query))
	scanner.Init(file, []byte(query), nil, token.ScanComments)

	// run the scanner, fill the output slice
	for {
		_, tok, lit := scanner.Scan()
		if tok == token.EOF {
			break
		}
		if lit == "\n" {
			continue
		}

		lexx = append(lexx, Lexemma{Token: tok, Litera: lit})
	}

	return lexx
}
