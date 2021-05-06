package lexemma

import "github.com/seggga/csvquery/token"

type Lexemma struct {
	Token  token.Token
	Litera string
}

//
func IsOperator(l *Lexemma) bool {
	switch l.Litera {
	case ">", ">=", "<", "<=", "==", "=":
		return true
	case "AND", "OR":
		return true
	default:
		return false

	}
}

func IsOperand(l *Lexemma) bool {

	if IsOperator(l) {
		return false
	}

	switch l.Litera {
	case "(", ")":
		return false
	default:
		return true
	}
}
