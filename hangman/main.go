package main

import (
	"fmt"
	"math/rand"
	"time"
)

var dictionary = []string{
	"Zombie",
	"Gopher",
	"United States of America",
	"Turkey",
	"Istanbul",
	"Hangman",
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {

		fmt.Println(getRandomWord())
	}
}

func getRandomWord() string {
	targetWord := dictionary[rand.Intn(len(dictionary))]

	return targetWord
}
