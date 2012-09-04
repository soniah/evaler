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
)

// vim: tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab tw=72
