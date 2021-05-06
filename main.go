package main

import (
	"fmt"

	"github.com/seggga/csvquery/token"
)

func main() {

	// query := []string{"age", ">", "40", "AND", "(", "city", "=", "Tokyo", "OR", "new", ">=", "1000", ")"}
	// expected := []string{"age", "40", ">", "city", "Tokyo", "=", "new", "1000", ">=", "OR", "AND"}
	// got := rpn.ConvertToRPN(query)

	// fmt.Printf("expected: %v\n", expected)
	// fmt.Printf("     got: %v\n", got)

	query := "humpty_dumpty >= set (on the wall) AND <= 10 times had a great fall OR (not=yes)"
	expected := []string{"humpty_dumpty", ">=", "set", "(", "on", "the", "wall", ")", "AND", "<=", "10", "times", "had", "a", "great", "fall", "OR", "(", "not", "=", "yes", ")"}

	got := token.SplitToTokens(query)

	fmt.Printf("got wrong slice length:\nexpected %d: %v\n     got %d: %v", len(expected), expected, len(got), got)

}
