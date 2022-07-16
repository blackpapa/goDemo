package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Andy")
	fmt.Println(message)
}
