package main

import (
	"fmt"
	"math"
)

func main() {
	var a int = 50
	var b string = "I love Go"

	a = 100
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	const c = 55
	//c = 89			// 不允许重新赋值
	fmt.Println("c =", c)

	// 常量的值会在编译的时候确定。因为函数调用发生在运行时，
	// 所以不能将函数的返回值赋值给常量
	fmt.Println("Hello, playground")
	var d1 = math.Sqrt(4) // 允许
	//const d2 = math.Sqrt(4)	// 不允许
	fmt.Println("d1 =", d1)

	const hello = "Hello World"
	fmt.Printf("type of hello is %T\n", hello)

	var name = "Sam"
	fmt.Printf("type %T value %v\n", name, name)

	const typedHello string = "Hello World"
	fmt.Printf("type %T value %v\n", typedHello, typedHello)

	// Go 是一个强类型的语言，在分配过程中混合类型是不允许的
	var defaultName = "Sam" // 允许
	type myString string
	var customName myString = "Sam2" // 允许
	//customName = defaultName				// 不允许
	customName = myString(defaultName)
	fmt.Println("customName =", customName)

	const trueConst = true
	type myBool bool
	var defaultBool = trueConst       // 允许
	var customBool myBool = trueConst // 允许
	//defaultBool = customBool				// 不允许
	defaultBool = bool(customBool)
	fmt.Println("defaultBool =", defaultBool)

	// 数字常量
	const aa = 5
	var intVar int = aa
	var int32Var int32 = aa
	var float64Var float64 = aa
	var complex64Var complex64 = aa
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)

	var i = 5
	var j = 5.6
	var k = 5 + 6i
	fmt.Printf("i's type %T, j's type %T, c's type %T\n", i, j, k)

	// 数字表达式
	var m = 5.9 / 8
	fmt.Printf("a's type %T value %v", m, m)

}
