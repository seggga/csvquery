package parse

import (
	"testing"
)

func TestCheckQuery(t *testing.T) {

	query := `SELECT a, b, c FROM file.csv WHERE a > b AND b >= c`
	got := CheckQuery(query)

	if got != nil {
		t.Errorf("expected no errors, got %v", got)
	}
}
