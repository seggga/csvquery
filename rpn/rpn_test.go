package rpn

import "testing"

func TestConvertToRPN(t *testing.T) {
	query := []string{"age", ">", "40", "AND", "(", "city", "=", "Tokyo", "OR", "new", ">=", "1000", ")"}
	expected := []string{"city", "Tokyo", "=", "new", "1000", ">=", "OR", "age", "40", ">", "AND"}
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
