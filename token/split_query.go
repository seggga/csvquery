package token

func SplitQuery(query string) []Lexemma {

	var scanner Scanner
	var lexx []Lexemma

	fset := NewFileSet()
	file := fset.AddFile("", fset.Base(), len(query))
	scanner.Init(file, []byte(query), nil, ScanComments)

	// run the scanner, fill the output slice
	for {
		_, tok, lit := scanner.Scan()
		if tok == EOF {
			break
		}
		if lit == "\n" {
			continue
		}

		// cut double quotes from 'STRING' literal (for example "Tokyo", where quotes are a part of the 'lit')
		if tok.String() == "STRING" {
			lit = lit[1 : len(lit)-1]
		}
		lexx = append(lexx, Lexemma{Token: tok.String(), Litera: lit})
	}

	return lexx
}
