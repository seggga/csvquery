package rpn

func ConvertToRPN(query []string) []string {

	var rpn, stack []string
	var stackPriority, tokenPriority int

	for _, token := range query {
		// операнд сразу заносится в выходную строку
		if isOperand(token) {
			rpn = append(rpn, token)
			continue
		}

		// открывающая скобка сразу идет в стек
		if token == "(" {
			stack = append(stack, token)
			continue
		}

		// закрывающая скобка извлекает все элементы из стека до символа "("
		if token == ")" {
			for i := len(stack) - 1; i >= 0; i -= 1 {
				stackToken := stack[i]
				stack = stack[:i]

				if stackToken == "(" {
					break
				}

				rpn = append(rpn, stackToken)
			}
			continue
		}

		// оператор участвует в ветвлении
		if isOperator(token) {
			tokenPriority = getPriority(token)

			// стек пуст - записываем знак операции в стек
			if len(stack) == 0 {
				stack = append(stack, token)
				continue
			}

			// стек не пуст, извлекаем все операторы из стека, чей приоритет >= приоритету токена
			// затем помещаем токен в стек
			for i := len(stack) - 1; i >= 0; i -= 1 {

				stackToken := stack[i]
				stackPriority = getPriority(stackToken)

				// приоритет у токена больше, чем у вершины стека - кладем токен в стек
				if tokenPriority > stackPriority {
					stack = append(stack, token)
					break
				}

				stack = stack[:i]

				rpn = append(rpn, stackToken)
			}

			// если из стека извлекли все операторы, то токен заносим в стек
			if len(stack) == 0 {
				stack = append(stack, token)
			}

		} // if isOperator(token)
	} // for _, token := range query

	// все токены просмотрены, надо опустошить стек
	for i := len(stack) - 1; i >= 0; i -= 1 {
		stackToken := stack[i]
		stack = stack[:i]
		rpn = append(rpn, stackToken)
	}

	return rpn
}

func isOperator(token string) bool {

	switch token {
	case ">", ">=", "<", "<=", "==", "=":
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
	case "AND", "OR":
		return 2
	case ">", "<", "==", ">=", "<=", "=":
		return 3
	default:
		return 100 // never gonna be returned
	}
}
