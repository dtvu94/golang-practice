package main

import (
	"fmt"
	"log"

	"rsc.io/quote"

	"example.com/greetings"
)

func main() {
	fmt.Println("Hello world!")

	fmt.Println(quote.Go())

	// Get a greeting message and print it.
	message, err := greetings.Hello("Success")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	messageFailed, errFailed := greetings.Hello("")

	if errFailed != nil {
		log.Fatal(errFailed)
	}

	fmt.Println(messageFailed)
}
