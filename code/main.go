package main

import (
	"fmt"
)

func main()  {
	a := []int{1,2,5, 6}
	fmt.Println(reverse2(a))
}

func reverse(arr []int) []int {
	n := len(arr)
	var tmp int

	end := n - 1
	for start := 0; start <= end; start++ {
		tmp = arr[start]
		arr[start] = arr[end]
		arr[end] = tmp
		end--
	}

	return arr
}

func reverse2(arr []int) []int {
	n := len(arr)
	tmp := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		tmp[n-1-i] = arr[i]
	}
	return tmp
}


