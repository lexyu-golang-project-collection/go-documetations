package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Lex")
	fmt.Println(message)
}
