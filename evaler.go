// Package evaler implements a simple fp arithmetic expression evaluator.
//
// Evaler uses Dijkstra's Shunting Yard algorithm [1] to convert an infix
// expression to postfix/RPN format [2], then evaluates the RPN expression.
//
// The operators supported are: + - * / ** parentheses () and < >
//
// < (less than) and > (greater than) will get lowest precedence, all
// other precedence is as expected (BODMAS).
//
// < and > tests will evaluate to 0.0 for false and 1.0 for true, allowing
// expressions like:
// 3 * (1 < 2) -> 3.0
// 3 * (1 > 2) -> 0.0
//
// [1] http://en.wikipedia.org/wiki/Shunting-yard_algorithm
// [2] http://en.wikipedia.org/wiki/Reverse_Polish_notation

package evaler

import (
	"fmt"
	"github.com/soniah/evaler/stack"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var whitespace_rx = regexp.MustCompile(`\s+`)
var fp_rx = regexp.MustCompile(`(\d+(?:\.\d)?)`) // simple fp number
var operators = "-+**/<>"

// prec returns the operator's precedence
func prec(op string) (result int) {
	if op == "-" || op == "+" {
		result = 1
	} else if op == "*" || op == "/" {
		result = 2
	} else if op == "**" {
		result = 3
	}
	return
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
	return fp_rx.MatchString(token)
}

// convert2postfix converts an infix expression to postfix
func convert2postfix(tokens []string) []string {
	var stack stack.Stack
	var result []string
	for _, token := range tokens {

		if isOperator(token) {

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
			result = append(result, token)
		}

	}

	for !stack.IsEmpty() {
		pop, _ := stack.Pop()
		result = append(result, pop.(string))
	}

	return result
}

// evaluatePostfix takes a postfix expression and evaluates it
func evaluatePostfix(postfix []string) float64 {
	var stack stack.Stack
	var result float64
	var fp float64
	for _, token := range postfix {
		if isOperand(token) {
			fp, _ = strconv.ParseFloat(token, 64)
			stack.Push(fp)
		} else if isOperator(token) {
			op2, _ := stack.Pop()
			op1, _ := stack.Pop()
			switch token {
			case "**":
				result = math.Pow(op1.(float64), op2.(float64))
				stack.Push(result)
			case "*":
				result = op1.(float64) * op2.(float64)
				stack.Push(result)
			case "/":
				result = op1.(float64) / op2.(float64)
				stack.Push(result)
			case "+":
				result = op1.(float64) + op2.(float64)
				stack.Push(result)
			case "-":
				result = op1.(float64) - op2.(float64)
				stack.Push(result)
			case "<":
				if op1.(float64) < op2.(float64) {
					stack.Push(1.0)
				} else {
					stack.Push(0.0)
				}
			case ">":
				if op1.(float64) > op2.(float64) {
					stack.Push(1.0)
				} else {
					stack.Push(0.0)
				}
			}
		} else {
			panic("Error")
		}
	}
	retval, _ := stack.Pop()
	return retval.(float64)
}

// tokenise takes an expr string and converts it to a slice of tokens
//
// tokenise puts spaces around all non-numbers, removes leading and
// trailing spaces, then splits on spaces
//
func tokenise(expr string) []string {
	spaced := fp_rx.ReplaceAllString(expr, " ${1} ")
	symbols := []string{"(", ")"}
	for _, symbol := range symbols {
		spaced = strings.Replace(spaced, symbol, fmt.Sprintf(" %s ", symbol), -1)
	}
	stripped := whitespace_rx.ReplaceAllString(strings.TrimSpace(spaced), "|")
	return strings.Split(stripped, "|")
}

// Eval takes an infix string arithmetic expression, and evaluates it
//
// Usage:
//   result, err := evaler.Eval("1+2")
// Returns: the result of the evaluation, and any errors
//
func Eval(expr string) (result float64, err error) {
	defer func() {
		if e := recover(); e != nil {
			result = float64(0.0)
			err = fmt.Errorf("Invalid Expression: %s", expr)
		}
	}()

	tokens := tokenise(expr)
	postfix := convert2postfix(tokens)
	result = evaluatePostfix(postfix)
	if math.IsInf(result, 0) {
		return result, fmt.Errorf("Divide by Zero: %s", expr)
	}
	return result, nil
}

// vim: tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab tw=74
