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
	guessedLetters := initializeGuessedWords(targetWord)
	hangmanState := 0

	for !isGameOver(targetWord, guessedLetters, hangmanState) {
		renderGameState(targetWord, guessedLetters, hangmanState)

		input := readInput()

		fmt.Println(len(input))

		if len(input) != 1 {
			fmt.Println("Please provide a letter. Not number, not special characters, not words.")

			continue
		}

		letter := rune(input[0])

		if alreadyGuessed(guessedLetters, letter) {
			fmt.Printf("You already guessed the letter %c.\n", letter)

			continue
		}

		if isCorrectGuess(targetWord, letter) {
			guessedLetters[letter] = true
		} else {
			hangmanState++
		}
	}

	fmt.Println("Game Over")
	if isWordGuessed(targetWord, guessedLetters) {
		fmt.Println("You win")
	} else {
		fmt.Println("You lose")
	}
}

func alreadyGuessed(guessedLetters map[rune]bool, input rune) bool {
	return guessedLetters[input] == true
}

func isGameOver(targetWord string, guessedLetters map[rune]bool, hangmanState int) bool {
	return isWordGuessed(targetWord, guessedLetters) || isHangmanComplete(hangmanState)
}

func isWordGuessed(targetWord string, guessedLetters map[rune]bool) bool {
	for _, letter := range targetWord {
		if !guessedLetters[unicode.ToLower(letter)] {
			return false
		}
	}

	return true
}

func isHangmanComplete(hangmanState int) bool {
	return hangmanState == 9
}

func isCorrectGuess(targetWord string, guess rune) bool {
	return strings.ContainsRune(targetWord, guess)
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
		} else if guessedLetters[unicode.ToLower(letter)] {
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
