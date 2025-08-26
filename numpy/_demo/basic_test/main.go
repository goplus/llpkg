package main

import (
	"fmt"

	"numpy"

	"github.com/goplus/lib/py"
	"github.com/goplus/lib/py/std"
)

func main() {
	fmt.Println("=== NumPy Basic Test Demo ===")

	// 测试1: 创建数组
	fmt.Println("\n1. 创建数组测试:")

	// 创建一个简单的数组
	arr1 := py.List(1.0, 2.0, 3.0, 4.0, 5.0)
	arr2 := py.List(2.0, 4.0, 6.0, 8.0, 10.0)

	// 使用numpy进行数组运算
	arr3 := numpy.Add(arr1, arr2)
	std.Print(py.Str("result:"), arr3)

	fmt.Println("\n=== 测试完成 ===")
}
