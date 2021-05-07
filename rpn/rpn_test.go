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

	// got empty slice
	if got == nil {
		t.Fatal("got nil-slice")
	}

	// got slice with different length
	if len(got) != len(expect) {
		fmt.Printf("expected: %v\n", expect)
		fmt.Printf("     got: %v\n", got)

		t.Fatal("slices differ in size")
	}

	// slices with different elements
	for i := 0; i < len(expect); i += 1 {
		if got[i] != expect[i] {
			t.Fatal("slices differ in elements")
		}
	}
}

func TestInsertValues(t *testing.T) {

	initialRPN := []token.Lexemma{
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
	csvData := map[string]string{
		"age":       "35",
		"city":      "Tokyo",
		"new":       "1001",
		"some_data": "data",
	}

	expect := []token.Lexemma{
		{Token: "IDENT", Litera: "35"},
		{Token: "INT", Litera: "40"},
		{Token: "COMP", Litera: ">"},
		{Token: "IDENT", Litera: "Tokyo"},
		{Token: "STRING", Litera: "Tokyo"},
		{Token: "COMP", Litera: "="},
		{Token: "IDENT", Litera: "1001"},
		{Token: "INT", Litera: "1000"},
		{Token: "COMP", Litera: ">="},
		{Token: "OR", Litera: "OR"},
		{Token: "AND", Litera: "AND"},
	}

	got := InsertValues(csvData, initialRPN)

	// got empty slice
	if got == nil {
		t.Fatal("got nil-slice")
	}

	// got slice with different length
	if len(got) != len(expect) {
		fmt.Printf("expected: %v\n", expect)
		fmt.Printf("     got: %v\n", got)

		t.Fatal("slices differ in size")
	}

	// slices with different elements
	for i := 0; i < len(expect); i += 1 {
		if got[i] != expect[i] {
			t.Fatal("slices differ in elements")
		}
	}

}

func TestCalculateRPN(t *testing.T) {

	rpn := []token.Lexemma{
		{Token: "IDENT", Litera: "4"},
		{Token: "INT", Litera: "40"},
		{Token: "COMP", Litera: ">"},
		{Token: "IDENT", Litera: "Moscow"},
		{Token: "STRING", Litera: "Tokyo"},
		{Token: "COMP", Litera: "=="},
		{Token: "IDENT", Litera: "100"},
		{Token: "INT", Litera: "1000"},
		{Token: "COMP", Litera: ">="},
		{Token: "OR", Litera: "OR"},
		{Token: "AND", Litera: "AND"},
	}

	expect := false

	got, _ := CalculateRPN(rpn)

	if got != expect {
		t.Fatalf("the answer is wrong:\nexpected: %t\n     got: %t", expect, got)
	}
}
