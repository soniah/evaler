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

```+ - * / ** () < >```

`<` (less than) and `>` (greater than) get lowest precedence, all other
precedence is as expected -
[BODMAS](http://www.mathsisfun.com/operation-order-bodmas.html).

`<` and `>` tests will evaluate to 0.0 for false and 1.0 for true, allowing
expressions like:

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

The trigonometric operators code was written by `dem-waffles
<dem-waffles@server.fake>` and manually integrated by me - thank you.

Issues
------

The `math/big` library doesn't have an exponent function `**` and implenting one
for `big.Rat` numbers is non-trivial. As a work around, arguments are converted
to float64's, the calculation is done using the `math.Pow()` function, the
result is converted to a `big.Rat` and placed back on the stack.

Documentation
-------------

http://godoc.org/github.com/soniah/evaler

There are also a number of utility functions e.g. `BigratToFloat()`,
`BigratToInt()` that may be useful when working with evaler.

Author
------

Sonia Hamilton sonia@snowfrog.net

Dem Waffles dem-waffles@server.fake - trigonometric operators


License
-------

Modified BSD License (BSD-3)

