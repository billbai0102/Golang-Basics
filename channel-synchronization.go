package main

import (
	"fmt"
	"time"
)

func worker(done chan bool){
	fmt.Print("working... ")
	time.Sleep(time.Second / 2)
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	fmt.Println(<- done)
	fmt.Println(len(done))
}
