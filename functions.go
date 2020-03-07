package main

import "fmt"

func plus(a int, b int) int{
	return a + b
}

func plus3(a int, b int, c int) int{
	return a + b + c
}

func main() {
	fmt.Printf("1 + 2 = %d\n", plus(1, 2))
	fmt.Printf("1 + 2 + 3 = %d\n", plus3(1, 2, 3))
}