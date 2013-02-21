package evaler_test

import (
	"github.com/soniah/evaler"
	"math/big"
	"testing"
)

var simpleArithTests = []struct {
	in  string
	out *big.Rat
	ok  bool
}{
	{"5 + 2", big.NewRat(7, 1), true},            // simple plus
	{"5 - 2", big.NewRat(3, 1), true},            // simple minus
	{"5 * 2", big.NewRat(10, 1), true},           // simple multiply
	{"5 / 2", big.NewRat(5, 2), true},            // simple divide
	{"U + U", nil, false},                        // letters 1
	{"2 + U", nil, false},                        // broken 1
	{"2 +  ", nil, false},                        // broken 2
	{"+ 2 - + * ", nil, false},                   // broken 3
	{"5.5+2*(3+1)", big.NewRat(27, 2), true},     // complex 1
	{"(((1+2.3)))", big.NewRat(33, 10), true},    // complex 2
	{"(1+(2))*(5-2.5)", big.NewRat(15, 2), true}, // complex 3
	{"3*(2<4)", big.NewRat(3, 1), true},          // less than
	{"3*(2>4)", new(big.Rat), true},              // greater than
	{"5 / 0", nil, false},                        // divide by zero
}

func TestSimpleArith(t *testing.T) {
	for i, test := range simpleArithTests {
		ret, err := evaler.Eval(test.in)
		if ret == nil && test.out == nil {
			// ok, do nothing
		} else if ret == nil || test.out == nil {
			t.Errorf("#%d: %s: unexpected nil result: %v vs %v", i, test.in, ret, test.out)
		} else if ret.Cmp(test.out) != 0 {
			t.Errorf("#%d: %s: bad result: got %v expected %v", i, test.in, ret, test.out)
		}
		if (err == nil) != test.ok {
			t.Errorf("#%d: %s: unexpected err result: %t vs %t", i, test.in, (err == nil), test.ok)
		}
	}
}

/*

func (s *MySuite) TestExponent1(c *C) {
	res, err := evaler.Eval("2 ** 3")
	c.Check(res, Equals, float64(8.0))
	c.Check(err, IsNil)
}

func (s *MySuite) TestExponent2(c *C) {
	res, err := evaler.Eval("9.0**0.5")
	c.Check(res, Equals, float64(3.0))
	c.Check(err, IsNil)
}

*/
