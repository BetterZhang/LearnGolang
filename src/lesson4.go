package main

import "fmt"

func functionname() {
	// 译注: 表示这个函数不需要输入参数，且没有返回值
}

func calculateBill(price int, no int) int {
	var totalPrice = price * no // 商品总价 = 商品单价 * 数量
	return totalPrice
}

func calculateBill2(price, no int) int {
	var totalPrice = price * no // 商品总价 = 商品单价 * 数量
	return totalPrice
}

// 多返回值
func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter

}

// 命名返回值
func rectProps2(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return
}

/*
	函数
*/
func main() {
	price, no := 90, 6
	totalPrice := calculateBill2(price, no)
	fmt.Println("Total price is ", totalPrice)

	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area %f, Perimeter %f\n", area, perimeter)

	area2, perimeter2 := rectProps2(10.8, 5.6)
	fmt.Printf("Area %f, Perimeter %f\n", area2, perimeter2)

	// 空白符
	area3, _ := rectProps(10.8, 5.6) // 返回值周长被丢弃
	fmt.Printf("Area %f\n", area3)

}
