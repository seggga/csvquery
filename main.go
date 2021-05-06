package main

import (
	"fmt"

	"github.com/seggga/csvquery/lexemma"
)

func main() {

	// query := []string{"age", ">", "40", "AND", "(", "city", "=", "Tokyo", "OR", "new", ">=", "1000", ")"}
	// expected := []string{"age", "40", ">", "city", "Tokyo", "=", "new", "1000", ">=", "OR", "AND"}
	// got := rpn.ConvertToRPN(query)

	// fmt.Printf("expected: %v\n", expected)
	// fmt.Printf("     got: %v\n", got)

	query := `age > 40 AND (city_name == "Tokyo" OR new_issues <= 1000)`

	fmt.Println(query)
	fmt.Printf("%+v\n", lexemma.SplitQuery(query))

}
