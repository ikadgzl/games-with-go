package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
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

var inputReader *bufio.Reader = bufio.NewReader(os.Stdin)

func main() {
	rand.Seed(time.Now().UnixNano())

	targetWord := getRandomWord()
	guessed := initializeGuessedWords(targetWord)
	hangmanState := 0

	for {
		renderGameState(targetWord, guessed, hangmanState)

		input := readInput()

		fmt.Println(len(input))

		if len(input) != 1 {
			fmt.Println("Please provide a letter. Not number, not special characters, not words.")

			continue
		} else {

		}
	}

}

func renderGameState(targetWord string, guessedLetters map[rune]bool, hangmanState int) {
	fmt.Println(renderGuessingProgress(targetWord, guessedLetters))
	fmt.Println()
	fmt.Println(renderHangman(hangmanState))
}

func renderHangman(state int) string {
	data, err := ioutil.ReadFile(fmt.Sprintf("states/hangman%d", state))

	if err != nil {
		panic(err)
	}

	return string(data)
}

func renderGuessingProgress(targetWord string, guessedLetters map[rune]bool) string {
	var result string
	for _, letter := range targetWord {
		if letter == ' ' {
			result += "        "
		} else if guessedLetters[unicode.ToLower(letter)] == true {
			result += fmt.Sprintf("%c", letter)
		} else {
			result += " _ "
		}
	}

	return result
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

func readInput() string {
	fmt.Print(">>>  ")
	input, err := inputReader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(input)
}
