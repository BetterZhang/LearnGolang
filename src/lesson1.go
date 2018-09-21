package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello World")

	var age int // 变量声明
	fmt.Println("My age is", age)
	age = 29 // 赋值
	fmt.Println("My age is", age)
	age = 54 // 赋值
	fmt.Println("My new age is", age)

	var age2 = 29 // 声明变量并初始化
	fmt.Println("My age2 is", age2)

	var age3 = 29 // 可以推断类型
	fmt.Println("My age3 is", age3)

	// 声明多个变量
	var width, height = 100, 50
	fmt.Println("Width is", width, "height is", height)

	var width2, height2 int
	fmt.Println("Width is", width2, "height is", height2)
	width2 = 100
	height2 = 50
	width2, height2 = 110, 60
	fmt.Println("New width is", width2, "new height is", height2)

	// 声明不同类型的变量
	var (
		name4   = "naveen"
		age4    = 29
		height4 int
	)
	fmt.Println("My name is", name4, ", age is", age4, "and height is", height4)

	// 简短声明
	name5, age5 := "naveen", 29
	fmt.Println("My name is", name5, "age is", age5)

	a, b := 20, 30 // 声明变量a和b
	fmt.Println("a is", a, "b is", b)
	b, c := 40, 50 // b已经声明，但c尚未声明
	fmt.Println("b is", b, "c is", c)
	b, c = 80, 90 // 给已经声明的变量b和c赋新值
	fmt.Println("Changed b is", b, "c is", c)

	e, f := 145.8, 543.8
	g := math.Min(e, f)
	fmt.Println("minimum value is", g)

	age6 := 29 // age6是int类型
	//age6 = "naveen"		// 错误，尝试赋值一个字符串给int类型变量
	fmt.Println("My age6 is", age6)
}
