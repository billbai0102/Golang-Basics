package main

import "fmt"

func passByVal(val int){
	val = 0
}

func passByPointer(ptr *int)  {
	*ptr = 5
}

func main() {
	var x int = 1
	fmt.Println("init: ", x)

	passByVal(x)
	fmt.Println("pval: ", x)

	passByPointer(&x)
	fmt.Println("pptr: ", x)

	fmt.Println("vptr: ", &x)
}