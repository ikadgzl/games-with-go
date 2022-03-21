package main

import (
	"fmt"
	"math/rand"
	"time"
	"unicode"
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

	targetWord := getRandomWord()

	targetWord = "United States of America"

	guessed := initializeGuessedWords(targetWord)

	renderGameState(targetWord, guessed)

	guessed['s'] = true

	renderGameState(targetWord, guessed)

}

func getRandomWord() string {
	targetWord := dictionary[rand.Intn(len(dictionary))]

	return targetWord
}

func initializeGuessedWords(targetWord string) map[rune]bool {
	guessedLetters := map[rune]bool{}

	guessedLetters[unicode.ToLower(rune(targetWord[0]))] = true
	guessedLetters[unicode.ToLower(rune(targetWord[len(targetWord)-1]))] = true

	return guessedLetters
}

func renderGameState(targetWord string, guessedLetters map[rune]bool) {

	for _, letter := range targetWord {
		if letter == ' ' {
			fmt.Print("        ")
		} else if guessedLetters[unicode.ToLower(letter)] == true {
			fmt.Printf("%c", letter)
		} else {
			fmt.Print(" _ ")
		}
	}

	fmt.Println()
}
