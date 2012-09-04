// Package evaler implements a simple fp arithmetic expression evaluator.
//
// Evaler uses Dijkstra's Shunting Yard algorithm [1] to convert an infix
// expression to postfix/RPN format [2], then evaluates the RPN expression.
//
// The operators supported are: + - * / and parentheses ().
//
// [1] http://en.wikipedia.org/wiki/Shunting-yard_algorithm
// [2] http://en.wikipedia.org/wiki/Reverse_Polish_notation

package evaler

import (
	"evaler/stack"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var whitespace = regexp.MustCompile(`\s+`)
var fp = regexp.MustCompile(`\d+(?:\.\d)?`) // simple fp number
var operators = "-+*/"

// prec returns the operator's precedence
func prec(op string) int {
	result := 0
	if op == "-" || op == "+" {
		result = 1
	} else if op == "*" || op == "/" {
		result = 2
	}
	return result
}

// opGTE returns true if op1's precedence is >= op2
func opGTE(op1, op2 string) bool {
	return prec(op1) >= prec(op2)
}

// isOperator returns true if token is an operator
func isOperator(token string) bool {
	return strings.Contains(operators, token)
}

// isOperand returns true if token is an operand
func isOperand(token string) bool {
	return fp.MatchString(token)
}

func convert2postfix(tokens []string) []string {
	var stack stack.Stack
	var result []string
	for _, token := range tokens {

		if isOperator(token) {
			//fmt.Printf("token %s is op\n", token)

		OPERATOR:
			for {
				top, err := stack.Top()
				if err == nil && top != "(" {
					if opGTE(top.(string), token) {
						pop, _ := stack.Pop()
						result = append(result, pop.(string))
					} else {
						break OPERATOR
					}
				}
				break OPERATOR
			}
			stack.Push(token)

		} else if token == "(" {
			stack.Push(token)

		} else if token == ")" {
		PAREN:
			for {
				top, err := stack.Top()
				if err == nil && top != "(" {
					pop, _ := stack.Pop()
					result = append(result, pop.(string))
				} else {
					stack.Pop() // pop off "("
					break PAREN
				}
			}

		} else if isOperand(token) {
			//fmt.Printf("token %s is operand\n", token)
			result = append(result, token)
		}

		//fmt.Printf("stack  is: %v\n", stack)
		//fmt.Printf("result is: %v\n\n", result)
	}

	for !stack.IsEmpty() {
		pop, _ := stack.Pop()
		result = append(result, pop.(string))
	}

	//fmt.Printf("stack  is: %v\n", stack)
	//fmt.Printf("result is: %v\n\n", result)
	return result
}

func evaluatePostfix(postfix []string) float64 {
	var stack stack.Stack
	var result float64
	var fp float64
	for _, token := range postfix {
		if isOperand(token) {
			fp, _ = strconv.ParseFloat(token, 64)
			stack.Push(fp)
		} else if isOperator(token) {
			op1, _ := stack.Pop()
			op2, _ := stack.Pop()
			switch token {
			case "*":
				result = op1.(float64) * op2.(float64)
				stack.Push(result)
			case "/":
				// TODO handle div by zero
				result = op1.(float64) / op2.(float64)
				stack.Push(result)
			case "+":
				result = op1.(float64) + op2.(float64)
				stack.Push(result)
			case "-":
				result = op1.(float64) / op2.(float64)
				stack.Push(result)
			}
		} else {
			fmt.Println("Error")
		}
		//fmt.Printf("stack: %v\n", stack)
	}
	retval, _ := stack.Pop()
	return retval.(float64)
}

func Eval(expr string) (float64, error) {
	fix_parens := strings.Replace(expr, "(", " ( ", -1)
	fix_parens = strings.Replace(fix_parens, ")", " ) ", -1)
	stripped := whitespace.ReplaceAllString(strings.TrimSpace(fix_parens), "|")
	tokens := strings.Split(stripped, "|")

	postfix := convert2postfix(tokens)
	result := evaluatePostfix(postfix)
	//fmt.Printf("Result is: %f\n", result)
	return result, nil
}

// vim: tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab tw=74
