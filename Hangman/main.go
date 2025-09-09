package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"slices"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func main() {
	// pick random word
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	sliceOfword := []string{"golang", "javaScript", "python", "java", "haskell", "rust", "zig"}
	word := sliceOfword[random.Intn(len(sliceOfword))]

	// max false trys is 6 create '_''_''_' as size of the random word
	// create scanner also to take input for the user
	attempts := 6
	currentWordState := initWordState(word)
	scanner := bufio.NewScanner(os.Stdin)
	gussedLetters := make(map[string]bool)

	// start the game
	fmt.Println("Welcome to hangman game...")
	time.Sleep(4 * time.Second)

	for attempts > 0 {
		clearConsole()
		// print the ascii code draw
		fmt.Println(asciiStatesDrawHangman[attempts])
		displayCurrentState(currentWordState)
		userInput := getUserInput(scanner)

		if !isValidInput(userInput) {
			fmt.Println("Invalid input  please inter single character")
			time.Sleep(3 * time.Second)
			continue
		} else if gussedLetters[userInput] {
			fmt.Println("You already guessed that letter .")
			time.Sleep(1 * time.Second)
			continue
		}

		isCorrectGuess := updateWordIfCorrect(word, currentWordState, userInput)
		isTheWordCompleted := isCompleteWord(currentWordState)
		gussedLetters[userInput] = true

		if isCorrectGuess && isTheWordCompleted {
			clearConsole()
			fmt.Println("You are the  winner the word is ", word)
			break
		} else if !isCorrectGuess {
			attempts--
		}
		if attempts < 0 {
			clearConsole()
			fmt.Println("The man Is dead")
			fmt.Println("the word was is ", word)
			break
		}

	}
}

func getUserInput(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func initWordState(word string) []string {
	currentWordState := make([]string, len(word))

	for i := range currentWordState {
		currentWordState[i] = "_"
	}
	return currentWordState
}

func isValidInput(input string) bool {
	if _, err := strconv.Atoi(input); err == nil || utf8.RuneCountInString(input) > 1 || input == "" {
		return false
	}
	return true
}

func displayCurrentState(currentWord []string) {
	fmt.Println("Current word state : ", strings.Join(currentWord, ""))
}

func updateWordIfCorrect(word string, currentWordState []string, letter string) bool {
	updateWordIfCorrect := false

	for i, char := range word {
		if string(char) == letter {
			currentWordState[i] = letter
			updateWordIfCorrect = true

		}
	}
	return updateWordIfCorrect
}

func isCompleteWord(word []string) bool {
	return !slices.Contains(word, "_")
}

func clearConsole() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
