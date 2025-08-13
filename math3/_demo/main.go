package main

import (
	"github.com/goplus/lib/py"
	"github.com/goplus/lib/py/std"
	"github.com/goplus/llpkg/math"
)

func main() {
	// 初始化 Python
	py.Initialize()
	defer py.Finalize()
	
	// 测试 math.Sqrt
	result1 := math.Sqrt(py.Float(4.0))
	std.Print(py.Str("sqrt(4) = "), result1)
	
	// 测试 math.Cos
	result2 := math.Cos(py.Float(0.0))
	std.Print(py.Str("cos(0) = "), result2)
	
	// 测试 math.Pow
	result3 := math.Pow(py.Float(2.0), py.Float(3.0))
	std.Print(py.Str("2^3 = "), result3)
	
	// 测试 math.Sin
	result4 := math.Sin(py.Float(1.57)) // 接近 pi/2
	std.Print(py.Str("sin(1.57) = "), result4)
}
