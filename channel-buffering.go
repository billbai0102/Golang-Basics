package main

import "fmt"

func main() {
	messages := make(chan string, 3)

	messages <- "1. buffered"
	messages <- "2. channel"
	messages <- "3. hurrah!"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)

	messages <- "4. ?"
	fmt.Println("size:",len(messages))
}