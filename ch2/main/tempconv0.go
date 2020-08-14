package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func main() {
	fmt.Printf("%v = %v\n", BoilingC, CToF(BoilingC))
	fmt.Printf("%v = %v\n", FreezingC, CToF(FreezingC))
	fmt.Printf("%v = %v\n", AbsoluteZeroC, CToF(AbsoluteZeroC))
}
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (c Fahrenheit) String() string { return fmt.Sprintf("%g°F", c) }
