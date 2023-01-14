package main

import "fmt"

func main()  {
	result := func1(10)

	fmt.Println("result:", result)
}

func func1(n int) int {
	var result = 1
	for i := 1; i <= n; i++ {
		result *= i
	}

	return result
}
