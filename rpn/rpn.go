package rpn

func ConvertToRPN(query []string) []string {

	var rpn, stack []string
	var stackPriority, tokenPriority int

	for _, token := range query {
		if isOperand(token) {
			rpn = append(rpn, token)
			continue
		}

		if token == "(" {
			stack = append(stack, token)
			continue
		}
		if token == ")" { // извлечение всех тоекнов из стека до "("
			for i := len(stack); i >= 0; i -= 1 {
				stackToken := stack[i]
				stack = stack[:i]

				if stackToken == "(" {
					break
				}

				rpn = append(rpn, stackToken)
			}
			continue
		}

		if isOperator(token) {
			tokenPriority = getPriority(token)

			// стек пуст - записываем знак операции в стек
			if len(stack) == 0 {
				stack = append(stack, token)
				stackPriority = tokenPriority
				continue
			}

			// стек не пуст, извлекаем все токены, чей приоритет больше либо равен текущему
			for i := len(stack); i >= 0; i -= 1 {

				stackToken := stack[i]
				if isOperator(stackToken) {
					stackPriority = getPriority(stackToken)
				}

				// приоритет у текущей операции больше, чем у стека - кладем токен в стек
				if stackPriority < tokenPriority {
					stack = append(stack, token)
					break
				}

				stack = stack[:i]

				rpn = append(rpn, stackToken)
			}

		}

	}
	return nil
}

func isOperator(token string) bool {

	switch token {
	case ">", ">=", "<", "<=", "==":
		return true
	case "AND", "OR":
		return true
	default:
		return false

	}
}

func isOperand(token string) bool {

	if isOperator(token) {
		return false
	}

	switch token {
	case "(", ")":
		return false
	default:
		return true
	}
}

func getPriority(token string) int {
	switch token {
	case "(":
		return 0
	case ")":
		return 1
	case ">", "<", "==", ">=", "<=":
		return 2
	case "AND", "OR":
		return 3
	default:
		return 100 // never gonna be returned
	}
}
