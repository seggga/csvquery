package main

import (
	"fmt"

	"github.com/seggga/csvquery/rpn"
	"github.com/seggga/csvquery/token"
)

func main() {

	query := `age > 40 AND (city_name == "Tokyo" OR new_issues <= 1000)`
	queryTokens := token.SplitQuery(query)
	queryTokens = rpn.ConvertToRPN(queryTokens)

	valuesMap := map[string]string{
		"age":        "30",
		"city_name":  "Moscow",
		"new_issues": "1000",
	}
	currentSlice := rpn.InsertValues(valuesMap, queryTokens)

	got, err := rpn.CalculateRPN(currentSlice)

	fmt.Println("got:", got, "error:", err)
}
