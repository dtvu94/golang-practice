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

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting messages for the names.
	messages, errs := greetings.Hellos(names)
	if errs != nil {
		log.Fatal(errs)
	}
	// If no error was returned, print the returned map of
	// messages to the console.
	fmt.Println(messages)

	// Get a greeting message and print it.
	message, err := greetings.Hello("Success")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	messageFailed, errFailed := greetings.Hello("")

	if errFailed != nil {
		log.Fatal(errFailed)
	} else {
		fmt.Println(messageFailed)
	}
}
