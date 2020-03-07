package main

import "fmt"

func return2(a1 int, b1 int, a2 int, b2 int) (int, int) {
	return a1 + b1, a2 + b2
}

func main() {
	a, b := return2(1, 2, 3, 4)
	fmt.Println(a)
	fmt.Println(b)
}
