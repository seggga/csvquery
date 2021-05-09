package parse

import (
	"errors"
	"regexp"
	"strings"
)

var (
	pattern string = `^select\b.+\bfrom\b.+\bwhere\b.+`
	error1  error  = errors.New("SELECT statement should be on the first place in the query")
	error2  error  = errors.New("the query should be like 'SELECT something FROM somewhere WHERE conditions'")
)

// CheckQuery checks the user's query for matching the pattern "SELECT-FROM-WHERE"
func CheckQuery(query string) error {

	query = strings.ToLower(query)
	query = strings.TrimSpace(query)

	// check SELECT existance
	if !strings.HasPrefix(query, "select") {
		return error1
	}

	// check matching to the pattern
	matched, _ := regexp.Match(pattern, []byte(query))
	if !matched {
		return error2
	}

	return nil
}
