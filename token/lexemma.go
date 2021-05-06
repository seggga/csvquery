package token

type Lexemma struct {
	Token  string
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
