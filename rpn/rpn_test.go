package rpn

import (
	"fmt"
	"testing"

	"github.com/seggga/csvquery/token"
)

func TestConvertToRPN(t *testing.T) {
	query := []token.Lexemma{
		{Token: "IDENT", Litera: "age"},
		{Token: "COMP", Litera: ">"},
		{Token: "INT", Litera: "40"},
		{Token: "AND", Litera: "AND"},
		{Token: "PAREN", Litera: "("},
		{Token: "IDENT", Litera: "city"},
		{Token: "COMP", Litera: "="},
		{Token: "STRING", Litera: "Tokyo"},
		{Token: "OR", Litera: "OR"},
		{Token: "IDENT", Litera: "new"},
		{Token: "COMP", Litera: ">="},
		{Token: "INT", Litera: "1000"},
		{Token: "PAREN", Litera: ")"},
	}

	expect := []token.Lexemma{
		{Token: "IDENT", Litera: "age"},
		{Token: "INT", Litera: "40"},
		{Token: "COMP", Litera: ">"},
		{Token: "IDENT", Litera: "city"},
		{Token: "STRING", Litera: "Tokyo"},
		{Token: "COMP", Litera: "="},
		{Token: "IDENT", Litera: "new"},
		{Token: "INT", Litera: "1000"},
		{Token: "COMP", Litera: ">="},
		{Token: "OR", Litera: "OR"},
		{Token: "AND", Litera: "AND"},
	}

	got := ConvertToRPN(query)

	if got == nil {
		t.Fatal("got nil-slice")
	}
	if len(got) != len(expect) {
		fmt.Printf("expected: %v\n", expect)
		fmt.Printf("     got: %v\n", got)

		t.Fatal("slices differ in size")
	}
	for i := 0; i < len(expect); i += 1 {
		if got[i] != expect[i] {
			t.Fatal("slices differ in elements")
		}
	}
}

// func TestIsOperator(t *testing.T) {
// 	testTable := []struct {
// 		in       string
// 		expected bool
// 	}{{">", true},
// 		{"AND", true},
// 		{"==", true},
// 		{"asdf", false},
// 	}

// 	for _, table := range testTable {
// 		if isOperator(table.in) != table.expected {
// 			t.FailNow()
// 		}
// 	}
// }

func TestInsertValues(t *testing.T) {

}
