package main

import (
	"fmt"

	"github.com/seggga/csvquery/lexemma"
	"github.com/seggga/csvquery/token"
)

func main() {

	// query := []string{"age", ">", "40", "AND", "(", "city", "=", "Tokyo", "OR", "new", ">=", "1000", ")"}
	// expected := []string{"age", "40", ">", "city", "Tokyo", "=", "new", "1000", ">=", "OR", "AND"}
	// got := rpn.ConvertToRPN(query)

	// fmt.Printf("expected: %v\n", expected)
	// fmt.Printf("     got: %v\n", got)

	query := `age > 40 AND (city_name == "Tokyo" OR new_issues <= 1000)`

	lexx := lexemma.SplitQuery(query)
	// fmt.Println(query)
	// fmt.Printf("%+v\n", lexx)

	if lexx[1].Token == token.COMP {
		fmt.Println("comp")
	}

	fmt.Println(lexx[1].Token.String())
}
