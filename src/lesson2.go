package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)
	c := a && b
	fmt.Println("c:", c)
	d := a || b
	fmt.Println("d:", d)

	var e int = 89
	f := 95
	fmt.Println("Value of e is", e, "and f is", f)
	fmt.Printf("Type of e is %T, size of e is %d", e, unsafe.Sizeof(e))   // e的类型和大小
	fmt.Printf("\nType of f is %T, size of f is %d", f, unsafe.Sizeof(f)) // f的类型和大小

	m, n := 5.67, 8.97
	fmt.Printf("\nType of m %T n %T", m, n)

	sum := m + n
	diff := m - n
	fmt.Println("\nsum", sum, "diff", diff)

	no1, no2 := 56, 89
	fmt.Println("sum", no1+no2, "diff", no1-no2)

	c1 := complex(5, 7)
	c2 := 8 + 27i
	cAdd := c1 + c2
	fmt.Printf("type of c1 is %T\n", c1)
	fmt.Printf("type of c2 is %T\n", c2)
	fmt.Println("sum:", cAdd)
	cMul := c1 * c2
	fmt.Println("product:", cMul)

	first := "Naveen"
	last := "Ramanathan"
	name := first + " " + last
	fmt.Println("My name is", name)

	i := 55   // int
	j := 67.8 // float64
	//sum2 := i + j		// 不允许 int + float64
	sum2 := i + int(j) // j is converted to int
	fmt.Println("sum2:", sum2)

	k := 10
	var l float64 = float64(k) // 若没有显式转换，该语句会报错
	fmt.Println("l =", l)
}
