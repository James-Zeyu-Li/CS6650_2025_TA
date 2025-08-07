package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano())) // real random number generator
// var rng = rand.New(rand.NewSource(42)) // fixed seed for reproducibility

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name cannot be empty")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Hello, %v!",
		"Hi, %v! How are you?",
		"Greetings, %v!",
		"Welcome, %v!",
	}
	return formats[rng.Intn(len(formats))]
}

func main() {
	message, err := Hello("World")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(message)
}
