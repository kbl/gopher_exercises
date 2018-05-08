package mandelbrot

import (
	"fmt"
	"math/big"
)

var zero *big.Float = new(big.Float)

type ComplexFloat struct {
	r, i *big.Float
}

func NewComplexFloat() *ComplexFloat {
	cf := new(ComplexFloat)
	cf.r = new(big.Float)
	cf.i = new(big.Float)
	return cf
}

func (z *ComplexFloat) Add(x, y *ComplexFloat) *ComplexFloat {
	var r, i *big.Float = z.r, z.i

	r.Add(x.r, y.r)
	i.Add(x.i, y.i)

	return z
}

func (z *ComplexFloat) Mul(x, y *ComplexFloat) *ComplexFloat {
	var newR1 *big.Float = &big.Float{}
	var newR2 *big.Float = &big.Float{}
	var newI1 *big.Float = &big.Float{}
	var newI2 *big.Float = &big.Float{}

	var r, i *big.Float = z.r, z.i

	newR1 = newR1.Mul(x.r, y.r)
	newR2.Mul(x.i, y.i)
	r.Sub(newR1, newR2)

	newI1.Mul(x.r, y.i)
	newI2.Mul(x.i, y.r)
	i.Add(newI1, newI2)

	return z
}

func (c *ComplexFloat) Abs() *big.Float {
	var f1 *big.Float = &big.Float{}
	var f2 *big.Float = &big.Float{}

	f1.Mul(c.r, c.r)
	f2.Mul(c.i, c.i)
	f1.Add(f1, f2)

	return f1.Sqrt(f1)
}

func (c *ComplexFloat) String() string {
	var imaginative string
	var comparision int = c.i.Cmp(zero)

	if comparision == -1 { // negative
		imaginative = fmt.Sprintf("%gi", c.i)
	} else if comparision == 1 { // positive
		imaginative = fmt.Sprintf("+%gi", c.i)
	}

	if c.i.Signbit() {

	}
	return fmt.Sprintf("ComplexFloat(%v%s)", c.r, imaginative)
}
