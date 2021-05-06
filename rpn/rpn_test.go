package rpn

import "testing"

func TestConvertToRPN(t *testing.T) {
	query := []string{"age", ">", "40", "AND", "(", "city", "=", "Tokyo", "OR", "new", ">=", "1000", ")"}
	expected := []string{"age", "40", ">", "city", "Tokyo", "=", "new", "1000", ">=", "OR", "AND"}
	got := ConvertToRPN(query)

	if got == nil {
		t.Fatal("got nil-slice")
	}
	if len(got) != len(expected) {
		t.Fatal("slices differ in size")
	}
	for i := 0; i < len(expected); i += 1 {
		if got[i] != expected[i] {
			t.Fatal("slices differ in elements")
		}
	}
}

func TestIsOperator(t *testing.T) {
	testTable := []struct {
		in       string
		expected bool
	}{{">", true},
		{"AND", true},
		{"==", true},
		{"asdf", false},
	}

	for _, table := range testTable {
		if isOperator(table.in) != table.expected {
			t.FailNow()
		}
	}
}
