package token

type Lexema struct {
	Token  string
	Litera string
}

//
func IsOperator(l *Lexema) bool {
	switch l.Litera {
	case ">", ">=", "<", "<=", "==", "=":
		return true
	case "AND", "OR":
		return true
	default:
		return false

	}
}

func IsOperand(l *Lexema) bool {

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
