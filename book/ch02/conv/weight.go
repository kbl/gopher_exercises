package conv

import "fmt"

type Kilogram float64
type Pound float64

const weightRate = 0.45359237

func (k Kilogram) String() string { return fmt.Sprintf("%.2f lbs", k) }
func (k Kilogram) ToP() Pound     { return Pound(k / weightRate) }

func (p Pound) String() string { return fmt.Sprintf("%.2f kg", p) }
func (p Pound) ToK() Kilogram  { return Kilogram(p * weightRate) }
