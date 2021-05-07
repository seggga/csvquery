package rpn

import (
	"errors"

	"github.com/seggga/csvquery/token"
)

var (
	error1 error = errors.New("error in expression. Probably wrong query")
	error2 error = errors.New("the result was not obtained. Probably wrong query")
)

// CalculateRPN - implements the mathematical calculation of given slice
// represented as reverse polish notation (RPN)
func CalculateRPN(rpn []token.Lexemma) (bool, error) {

	for i := 0; i < len(rpn); {

		if !token.IsOperand(&rpn[i]) {
			i += 1
			continue
		}
		arg1 := rpn[i]

		if !token.IsOperand(&rpn[i+1]) {
			i += 1
			continue
		}
		arg2 := rpn[i+1]

		if !token.IsOperator(&rpn[i+2]) {
			i += 1
			continue
		}
		op := rpn[i+2]

		result, err := solveExpression(arg1, arg2, op)
		if err != nil {
			return false, err
		}

		// remove 3 elements from the rpn and move 'result' on that place
		newRPN := rpn[:i]
		newRPN = append(newRPN, result)
		newRPN = append(newRPN, rpn[i+3:]...)
		rpn = newRPN

		// start next iteration from '0' element in updated 'rpn'
		i = 0

		if len(rpn) == 1 {
			return rpn[0].Litera == "TRUE", nil
		}
	}

	return false, error2
}

// solveExpression calculates a given expression.
// As a result it makes token.Lexemma{Token: "IDENT", Litera: "TRUE"/"FALSE"}
// The main goal is to pass 'Litera' field with "TRUE" or "FALSE" value.
func solveExpression(arg1, arg2, op token.Lexemma) (token.Lexemma, error) {

	var result bool

	switch op.Litera {
	case "==":
		result = arg1.Litera == arg2.Litera
	case ">=":
		result = arg1.Litera >= arg2.Litera
	case "<=":
		result = arg1.Litera <= arg2.Litera
	case ">":
		result = arg1.Litera > arg2.Litera
	case "<":
		result = arg1.Litera < arg2.Litera
	case "AND":

		var operand1, operand2 bool

		// check if arg1 is 'true' / 'false'
		if arg1.Litera == "TRUE" {
			operand1 = true
		} else if arg1.Litera == "FALSE" {
			operand1 = false
		} else {
			return token.Lexemma{}, error1
		}

		// check if arg2 is 'true' / 'false'
		if arg2.Litera == "TRUE" {
			operand2 = true
		} else if arg2.Litera == "FALSE" {
			operand2 = false
		} else {
			return token.Lexemma{}, error1
		}

		result = operand1 && operand2

	case "OR":

		var operand1, operand2 bool

		// check if arg1 is 'true' / 'false'
		if arg1.Litera == "TRUE" {
			operand1 = true
		} else if arg1.Litera == "FALSE" {
			operand1 = false
		} else {
			return token.Lexemma{}, error1
		}

		// check if arg2 is 'true' / 'false'
		if arg2.Litera == "TRUE" {
			operand2 = true
		} else if arg2.Litera == "FALSE" {
			operand2 = false
		} else {
			return token.Lexemma{}, error1
		}

		result = operand1 || operand2
	}

	if result {
		return token.Lexemma{Token: "IDENT", Litera: "TRUE"}, nil
	}
	return token.Lexemma{Token: "IDENT", Litera: "FALSE"}, nil
}
