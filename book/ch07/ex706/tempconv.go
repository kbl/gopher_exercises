package ex706

import (
	"book/ch02/ex201"
	"flag"
	"fmt"
)

type celsiusFlag struct {
	ex201.Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celsius = ex201.Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = ex201.Fahrenheit(value).ToC()
		return nil
	case "K", "°K":
		f.Celsius = ex201.Kelvin(value).ToC()
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

var _ flag.Value = (*celsiusFlag)(nil)

func CelsiusFlag(name string, value ex201.Celsius, usage string) *ex201.Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
