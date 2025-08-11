package main

import (
	"github.com/goplus/lib/py"
	"github.com/goplus/llpkg/math"
)

func main() {
	// Test basic math functions
	pi := math.Pi()
	py.Print(pi)

	// Test trigonometric functions
	angle := py.Float(3.14159 / 4) // 45 degrees in radians
	sinVal := math.Sin(angle)
	cosVal := math.Cos(angle)

	py.Print(py.Str("sin(45°) = "))
	py.Print(sinVal)
	py.Print(py.Str("cos(45°) = "))
	py.Print(cosVal)

	// Test power function
	base := py.Float(2.0)
	exponent := py.Float(3.0)
	result := math.Pow(base, exponent)

	py.Print(py.Str("2^3 = "))
	py.Print(result)

	// Test square root
	sqrtInput := py.Float(16.0)
	sqrtResult := math.Sqrt(sqrtInput)

	py.Print(py.Str("sqrt(16) = "))
	py.Print(sqrtResult)
}
