package main

import (
	"fmt"
	"geometry/rectangle"
	_ "geometry/rectangle"
	"log"
)

/*
 * 1.包级别变量
 */
var rectLen, rectWidth float64 = 6, 7

// 建议在 import 语句下面的包级别范围中写上错误屏蔽器
//var _ = rectangle.Area		// 错误屏蔽器

/*
 * 2. init 函数会检查长和宽是否大于0
 */
func init() {
	fmt.Println("main package initialized")
	if rectLen < 0 {
		log.Fatal("length is less than zero")
	}
	if rectWidth < 0 {
		log.Fatal("width is less than zero")
	}
}

func init() {
	fmt.Println("another main package initialized")
}

func main() {
	//var rectLen, rectWidth float64 = 6, 7
	fmt.Println("Geometrical shape properties")
	fmt.Printf("Area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	fmt.Printf("Diagonal of the rectangle %.2f ", rectangle.Diagonal(rectLen, rectWidth))
}
