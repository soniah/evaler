evaler
======

Package evaler implements a simple fp arithmetic expression evaluator.

Evaler uses Dijkstra's Shunting Yard algorithm [1] to convert an infix
expression to postfix/RPN format [2], then evaluates the RPN expression. The
implementation is adapted from a Java implementation at [3]. The results are
returned as a `math/big *big.Rat`.

This is release 2.0. The previous version that returned results as float64
is in the branch float64.

Usage
-----

```go
result, err := evaler.Eval("1+2")
```

Operators
---------

The operators supported are:

```+ - * / ** () < >```

< (less than) and > (greater than) will get lowest precedence, all
other precedence is as expected (BODMAS [4]).

< and > tests will evaluate to 0.0 for false and 1.0 for true, allowing
expressions like:

```
3 * (1 < 2) # returns 3.0
3 * (1 > 2) # returns 0.0
```

Issues
------

The 2.0 release currently doesn't support exponents (** operator). This
is WIP (work in progress).

Author
------

Sonia Hamilton

http://www.snowfrog.net

sonia@snowfrog.net

License
-------

Modified BSD License (BSD-3)

Links
-----

[1] http://en.wikipedia.org/wiki/Shunting-yard_algorithm

[2] http://en.wikipedia.org/wiki/Reverse_Polish_notation

[3] http://willcode4beer.com/design.jsp?set=evalInfix

[4] http://www.mathsisfun.com/operation-order-bodmas.html

