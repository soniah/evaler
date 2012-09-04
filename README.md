evaler
======

Package evaler implements a simple fp arithmetic expression evaluator.

Evaler uses Dijkstra's Shunting Yard algorithm [1] to convert an infix
expression to postfix/RPN format [2], then evaluates the RPN expression. The
implementation is adapted from a Java implementation at [3].

The operators supported are: + - * / and parentheses ().

This is release 0.1 - error handling and testing still needs to be added.

[1] http://en.wikipedia.org/wiki/Shunting-yard_algorithm
[2] http://en.wikipedia.org/wiki/Reverse_Polish_notation
[3] http://willcode4beer.com/design.jsp?set=evalInfix

Author
------

Sonia Hamilton

http://www.snowfrog.net

sonia@snowfrog.net
