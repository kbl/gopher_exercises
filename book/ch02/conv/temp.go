package conv

import "fmt"

type Celsius float64
type Farenheit float64
type Kelvin float64

const (
    AbsoluteZeroC Celsius = -273.15
    FreezingC     Celsius = 0
    BoilingC      Celsius = 100
)

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
func (c Celsius) ToF() Farenheit { return Farenheit(c * 9 / 5 + 32) }
func (c Celsius) ToK() Kelvin { return Kelvin(c + 273.15) }

func (f Farenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (f Farenheit) ToC() Celsius { return Celsius((f - 32) * 5 / 9) }
func (f Farenheit) ToK() Kelvin { return f.ToC().ToK() }

func (k Kelvin) String() string { return fmt.Sprintf("%g°K", k) }
func (k Kelvin) ToC() Celsius { return Celsius(k - 273.15) }
func (k Kelvin) ToF() Farenheit { return k.ToC().ToF() }
