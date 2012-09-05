package evaler_test

import (
	"github.com/soniah/evaler"
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestSimpleArithmeticPlus(c *C) {
	res, err := evaler.Eval("5 + 2")
	c.Check(res, Equals, float64(7))
	c.Check(err, IsNil)
}

func (s *MySuite) TestSimpleArithmeticMinus(c *C) {
	res, err := evaler.Eval("5 - 2")
	c.Check(res, Equals, float64(3))
	c.Check(err, IsNil)
}

func (s *MySuite) TestSimpleArithmeticTimes(c *C) {
	res, err := evaler.Eval("5 * 2")
	c.Check(res, Equals, float64(10))
	c.Check(err, IsNil)
}

func (s *MySuite) TestSimpleArithmeticDivide(c *C) {
	res, err := evaler.Eval("5 / 2")
	c.Check(res, Equals, float64(2.5))
	c.Check(err, IsNil)
}

func (s *MySuite) TestDivideZero(c *C) {
	_, err := evaler.Eval("5 / 0")
	c.Check(err, ErrorMatches, "Divide by Zero:.*")
}

func (s *MySuite) TestLetters1(c *C) {
	res, err := evaler.Eval("U + U")
	c.Check(res, Equals, float64(0.0))
	c.Check(err, ErrorMatches, "Invalid Expression:.*")
}

func (s *MySuite) TestLetters2(c *C) {
	res, err := evaler.Eval("2 + U")
	c.Check(res, Equals, float64(0.0))
	c.Check(err, ErrorMatches, "Invalid Expression:.*")
}

func (s *MySuite) TestBrokenExpr1(c *C) {
	res, err := evaler.Eval("2 + ")
	c.Check(res, Equals, float64(0.0))
	c.Check(err, ErrorMatches, "Invalid Expression:.*")
}

func (s *MySuite) TestBrokenExpr2(c *C) {
	res, err := evaler.Eval("+ 2 - + * ")
	c.Check(res, Equals, float64(0.0))
	c.Check(err, ErrorMatches, "Invalid Expression:.*")
}

func (s *MySuite) TestComplex1(c *C) {
	res, err := evaler.Eval("5.5+2*(3+1)")
	c.Check(res, Equals, float64(13.5))
	c.Check(err, IsNil)
}

func (s *MySuite) TestComplex2(c *C) {
	res, err := evaler.Eval("(((1+2.3)))")
	c.Check(res, Equals, float64(3.3))
	c.Check(err, IsNil)
}

func (s *MySuite) TestComplex3(c *C) {
	res, err := evaler.Eval("(1+(2))*(5-2.5)")
	c.Check(res, Equals, float64(7.5))
	c.Check(err, IsNil)
}

// vim: tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab tw=74
