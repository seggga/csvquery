package token

func SplitToTokens(query string) []string {

	var scanner Scanner
	var tokens []string

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

		tokens = append(tokens, lit)
	}

	return tokens
}
