package conv

import "fmt"

type Meter float64
type Feet float64

const lengthRate = 0.3048

func (m Meter) String() string { return fmt.Sprintf("%.2f m", m) }
func (m Meter) ToF() Feet      { return Feet(m / lengthRate) }

func (f Feet) String() string { return fmt.Sprintf("%.2f ft", f) }
func (f Feet) ToM() Meter     { return Meter(f * lengthRate) }
