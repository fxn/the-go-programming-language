/*
   Exercise 2.1: Add types, constants, and functions to tempconv for processing
   temperatures in the Kelvin scale, where zero Kelvin is -273.15ºC and a
   difference of 1K has the same magnitude as 1ºC.
*/
package main

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%gºC", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%gºC", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }

func CtoF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FtoC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func CtoK(c Celsius) Kelvin     { return Kelvin(c - AbsoluteZeroC) }
func KtoC(k Kelvin) Celsius     { return Celsius(k) + AbsoluteZeroC }
func FtoK(f Fahrenheit) Kelvin  { return CtoK(FtoC(f)) }
func KtoF(k Kelvin) Fahrenheit  { return CtoF(KtoC(k)) }

func main() {
	fmt.Printf("0K = %gºC\n", KtoC(0))
	fmt.Printf("0ºC = %gK\n", CtoK(0))
	fmt.Printf("0ºF = %gK\n", FtoK(0))
}
