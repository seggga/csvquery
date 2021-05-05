package rpn

import "testing"

func ConvertToRPNTest(t *testing.T) {
	query := []string{"age", ">", "40", "AND", "(", "city", "=", "Tokyo", "OR", "new", ">=", "1000", ")"}
	expected := []string{"city", "Tokyo", "=", "new", "1000", ">=", "OR", "age", "40", ">", "AND"}
	got := ConvertToRPN(query)

	for i := 0; i < len(expected); i += 1 {
		if got[i] != expected[i] {
			t.Error("arrays differ")
		}
	}
}
