evaler
======

[![Build Status](https://travis-ci.org/soniah/evaler.svg?branch=master)](https://travis-ci.org/soniah/evaler)
[![Coverage](http://gocover.io/_badge/github.com/soniah/evaler)](http://gocover.io/github.com/soniah/evaler)
[![GoDoc](https://godoc.org/github.com/soniah/evaler?status.png)](http://godoc.org/github.com/soniah/evaler)
https://github.com/soniah/evaler

Package evaler implements a simple floating point arithmetic expression evaluator.

Evaler uses Dijkstra's [Shunting Yard algorithm](http://en.wikipedia.org/wiki/Shunting-yard_algorithm) to convert an
infix expression to [postfix/RPN format](http://en.wikipedia.org/wiki/Reverse_Polish_notation), then evaluates
the RPN expression. The implementation is adapted from a [Java implementation](http://willcode4beer.com/design.jsp?set=evalInfix). The results
are returned as a `*big.Rat`.

Usage
-----

```go
result, err := evaler.Eval("1+2")
```

Operators
---------

The operators supported are:

```+ - * / ^ ** () < > <= >= == !=```

(`^` and `**` are both exponent operators)

Logical operators like `<` (less than) or `>` (greater than) get lowest precedence,
all other precedence is as expected -
[BODMAS](http://www.mathsisfun.com/operation-order-bodmas.html).

Logical tests like `<` and `>` tests will evaluate to 0.0 for false and 1.0
for true, allowing expressions like:

```
3 * (1 < 2) # returns 3.0
3 * (1 > 2) # returns 0.0
```

Minus implements both binary and unary operations.

See `evaler_test.go` for more examples of using operators.

Trigonometric Operators
-----------------------

The trigonometric operators supported are:

```sin, cos, tan, ln, arcsin, arccos, arctan```

For example:

```
cos(1)
sin(2-1)
sin(1)+2**2
```

See `evaler_test.go` for more examples of using trigonometric operators.

Variables
---------

`EvalWithVariables()` allows variables to be passed into expressions,
for example evaluate `"x + 1"`, where `x=5`.

See `evaler_test.go` for more examples of using variables.

Issues
------

The `math/big` library doesn't have an exponent function `**` and implenting one
for `big.Rat` numbers is non-trivial. As a work around, arguments are converted
to float64's, the calculation is done using the `math.Pow()` function, the
result is converted to a `big.Rat` and placed back on the stack.

* floating point numbers missing leading digits (like `".5 * 2"`) are failing - PR's welcome

Documentation
-------------

http://godoc.org/github.com/soniah/evaler

There are also a number of utility functions e.g. `BigratToFloat()`,
`BigratToInt()` that may be useful when working with evaler.

Contributions
-------------

Contributions are welcome.

If you've never contributed to a Go project before here is an example workflow.

1. [fork this repo on the GitHub webpage](https://github.com/soniah/evaler/fork)
1. `go get github.com/soniah/evaler`
1. `cd $GOPATH/src/github.com/soniah/evaler`
1. `git remote rename origin upstream`
1. `git remote add origin git@github.com:<your-github-username>/evaler.git`
1. `git checkout -b development`
1. `git push -u origin development` (setup where you push to, check it works)

Author
------

Sonia Hamilton sonia@snowfrog.net

Dem Waffles dem-waffles@server.fake - trigonometric operators

License
-------

Modified BSD License (BSD-3)

Links
-----

[1] http://en.wikipedia.org/wiki/Shunting-yard_algorithm

[2] http://en.wikipedia.org/wiki/Reverse_Polish_notation

[3] http://willcode4beer.com/design.jsp?set=evalInfix

[4] http://www.mathsisfun.com/operation-order-bodmas.html

