package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func randomFormat() string {
	// A slice of message format
	formats := []string{
		"Hi, %v. Welcome",
		"Greet to see you, %v",
		"Hail, %v! Well met!",
	}

	//return a randomly selected message
	return formats[rand.Intn(len(formats))]
}

// Hero returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message
	if name == "" {
		return "", errors.New("name is empty")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}
