package main

import "fmt"

func main() {
	num := 10
	if num%2 == 0 {
		fmt.Println("the number is even")
	} else {
		fmt.Println("the number is odd")
	}

	// num2 只能从 if 和 else 中访问
	if num2 := 15; num2%2 == 0 {
		fmt.Println("the number is even")
	} else {
		fmt.Println("the number is odd")
	}

	num3 := 99
	if num3 <= 50 {
		fmt.Println("number is less than or equal to 50")
	} else if num3 >= 51 && num3 <= 100 {
		fmt.Println("number is between 51 and 100")
	} else {
		fmt.Println("number is greater than 100")
	}

}
