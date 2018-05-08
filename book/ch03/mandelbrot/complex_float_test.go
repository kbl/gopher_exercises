package mandelbrot

import (
	"fmt"
	"math/big"
	"math/cmplx"
	"testing"
)

func TestNewInstance(t *testing.T) {
	var f *ComplexFloat = NewComplexFloat()
	var zero *big.Float = new(big.Float)

	if f.r.Cmp(zero) != 0 {
		t.Errorf("%v != 0, wants %v", f.r, zero)
	}
	if f.i.Cmp(zero) != 0 {
		t.Errorf("%v != 0, wants %v", f.i, zero)
	}
}

func TestStringRepresentation(t *testing.T) {
	number := &ComplexFloat{
		r: big.NewFloat(1.5),
		i: big.NewFloat(2.5),
	}

	repr := fmt.Sprintf("%v", number)
	expected := "ComplexFloat(1.5+2.5i)"

	if repr != expected {
		t.Errorf("%s != %s", repr, expected)
	}

	number.i = big.NewFloat(-2.5)
	repr = fmt.Sprintf("%v", number)
	expected = "ComplexFloat(1.5-2.5i)"

	if repr != expected {
		t.Errorf("%s != %s", repr, expected)
	}

	number.i = big.NewFloat(0)
	repr = fmt.Sprintf("%v", number)
	expected = "ComplexFloat(1.5)"

	if repr != expected {
		t.Errorf("%s != %s", repr, expected)
	}
}

func TestAddition(t *testing.T) {
	var aComplex complex128 = 1 + 1i
	var bComplex complex128 = 2 + 2i

	result := aComplex + bComplex

	aFloat := &ComplexFloat{
		r: big.NewFloat(real(aComplex)),
		i: big.NewFloat(imag(aComplex)),
	}
	bFloat := &ComplexFloat{
		r: big.NewFloat(real(bComplex)),
		i: big.NewFloat(imag(bComplex)),
	}

	cFloat := NewComplexFloat()
	cFloat.Add(aFloat, bFloat)

	// shouldn't change arguments
	aRepr := fmt.Sprintf("%v", aFloat)
	expected := "ComplexFloat(1+1i)"
	if aRepr != expected {
		t.Errorf("%s != %s", aRepr, expected)
	}

	bRepr := fmt.Sprintf("%v", bFloat)
	expected = "ComplexFloat(2+2i)"
	if bRepr != expected {
		t.Errorf("%s != %s", bRepr, expected)
	}

	// addition operation
	r, _ := cFloat.r.Float64()
	i, _ := cFloat.i.Float64()
	if r != real(result) {
		t.Errorf("%v != %v", r, real(result))
	}
	if i != imag(result) {
		t.Errorf("%v != %v", i, imag(result))
	}
}

func TestMultiplication(t *testing.T) {
	var aComplex complex128 = 1 + 2i
	var bComplex complex128 = 3 + 4i

	result := aComplex * bComplex

	aFloat := &ComplexFloat{
		r: big.NewFloat(real(aComplex)),
		i: big.NewFloat(imag(aComplex)),
	}
	bFloat := &ComplexFloat{
		r: big.NewFloat(real(bComplex)),
		i: big.NewFloat(imag(bComplex)),
	}

	cFloat := NewComplexFloat()
	cFloat.Mul(aFloat, bFloat)

	// shouldn't change arguments
	aRepr := fmt.Sprintf("%v", aFloat)
	expected := "ComplexFloat(1+2i)"
	if aRepr != expected {
		t.Errorf("%s != %s", aRepr, expected)
	}

	bRepr := fmt.Sprintf("%v", bFloat)
	expected = "ComplexFloat(3+4i)"
	if bRepr != expected {
		t.Errorf("%s != %s", bRepr, expected)
	}

	// multiplication operation
	r, _ := cFloat.r.Float64()
	i, _ := cFloat.i.Float64()
	if r != real(result) {
		t.Errorf("%v != %v", r, real(result))
	}
	if i != imag(result) {
		t.Errorf("%v != %v", i, imag(result))
	}
}

func TestAbsolute(t *testing.T) {
	var builtIn complex128 = 3 + 4i

	expected := cmplx.Abs(builtIn)

	fmt.Println(expected)

	custom := &ComplexFloat{
		r: big.NewFloat(real(builtIn)),
		i: big.NewFloat(imag(builtIn)),
	}

	var abs *big.Float = custom.Abs()

	// shouldn't change arguments
	repr := fmt.Sprintf("%v", custom)
	expectedRepr := "ComplexFloat(3+4i)"
	if repr != expectedRepr {
		t.Errorf("%s != %s", repr, expectedRepr)
	}

	// absolute value operation
	result, _ := abs.Float64()
	if result != expected {
		t.Errorf("%v != %v", result, expected)
	}
}
