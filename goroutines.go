package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for x := 0; x < 3; x++{
		fmt.Println(from, ":", x)
	}
}

func main() {
	f("direct") // no async

	go f("goroutine") // async
	// time.Sleep(time.Nanosecond)

	go func(msg string){
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
