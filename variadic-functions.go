package main

import "fmt"

func sumAll(nums ...int) int{
	var total int = 0
	for x := 0; x < len(nums); x++{
		total += nums[x]
	}
	return total
}

func main() {
	fmt.Println(sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
