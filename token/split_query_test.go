package token

import "testing"

func TestSplitQuery(t *testing.T) {

	query := `age > 40 AND (city_name == "Tokyo" OR new_issues <= 1000)`

	expected := []Lexemma{
		{Token: "IDENT", Litera: "age"},
		{Token: "COMP", Litera: ">"},
		{Token: "INT", Litera: "40"},
		{Token: "AND", Litera: "AND"},
		{Token: "PAREN", Litera: "("},
		{Token: "IDENT", Litera: "city_name"},
		{Token: "COMP", Litera: "=="},
		{Token: "STRING", Litera: "Tokyo"},
		{Token: "OR", Litera: "OR"},
		{Token: "IDENT", Litera: "new_issues"},
		{Token: "COMP", Litera: "<="},
		{Token: "INT", Litera: "1000"},
		{Token: "PAREN", Litera: ")"},
	}

	got := SplitQuery(query)

	// got empty slice
	if got == nil {
		t.Fatal("got nil-slice")
	}

	// got slice with different length
	if len(got) != len(expected) {
		t.Fatalf("got wrong slice length:\nexpected %d: %v\n     got %d: %v", len(expected), expected, len(got), got)
	}

	// slices with different elements
	for i := range expected {
		if expected[i] != got[i] {
			t.Fatalf("slice mismatch on element %d: expected: %s\n     got: %s", i, expected[i], got[i])
		}
	}
}
