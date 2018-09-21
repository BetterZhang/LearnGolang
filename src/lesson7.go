package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}

	fmt.Println()

	for i := 1; i <= 10; i++ {
		if i > 5 {
			break
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\nline after for loop\n")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}

	fmt.Println()

	j := 0
	for j <= 10 {
		fmt.Printf("%d ", j)
		j += 2
	}

	fmt.Println()

	j1 := 0
	for j1 <= 10 {
		fmt.Printf("%d ", j1)
		j1 += 2
	}

	fmt.Println()

	for m, n := 10, 1; m <= 19 && n <= 10; m, n = m+1, n+1 {
		fmt.Printf("%d * %d = %d\n", m, n, m*n)
	}

	// 无限循环
	//for {
	//	fmt.Println("Hello World")
	//}
}
