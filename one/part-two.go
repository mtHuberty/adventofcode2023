package one

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var digitWords = []string{"one", "two", "three",
	"four", "five", "six",
	"seven", "eight", "nine"}

func SolvePartTwo() string {
	file, err := os.Open("one/input.txt")
	if err != nil {
		log.Fatal(fmt.Errorf("could not open file: %w", err))
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		firstDigitStr := getFirstDigit(line)
		lastDigitStr := getLastDigit(line)

		if _, err := strconv.Atoi(firstDigitStr); err != nil {
			fmt.Printf("could not convert first digit to int: %w", err)
			os.Exit(1)
		}
		if _, err := strconv.Atoi(lastDigitStr); err != nil {
			fmt.Printf("could not convert last digit to int: %w", err)
			os.Exit(1)
		}

		concatDigits := firstDigitStr + lastDigitStr
		sum, err := strconv.Atoi(concatDigits)
		if err != nil {
			fmt.Printf("could not convert concat digits to int: %w", err)
			os.Exit(1)
		}

		fmt.Println(concatDigits)

		total += sum
	}

	return strconv.Itoa(total)
}

func getFirstDigit(line string) string {
	possibleDigitWords := make([]string, len(digitWords))

	copy(possibleDigitWords, digitWords)
	indexForNextLetterOfDigitWord := 0

	for _, char := range line {

		// If we find a string representation of a number (1, 2, 3...) we found it
		if _, err := strconv.Atoi(string(char)); err == nil {
			return string(char)
		}

		// Otherwise we have to keep checking for ("one", "two"...)
		possibleDigitWords = getPossibleDigitWords(possibleDigitWords, char, indexForNextLetterOfDigitWord)

		if len(possibleDigitWords) == 0 {
			fmt.Println("found no words")
			copy(possibleDigitWords, digitWords)
			indexForNextLetterOfDigitWord = 0
			continue
		} else if len(possibleDigitWords) == 1 && len(possibleDigitWords[0]) == indexForNextLetterOfDigitWord {
			return possibleDigitWords[0]
		} else {
			indexForNextLetterOfDigitWord++
		}
	}

	fmt.Println("hit end of line without finding a digit")
	os.Exit(1)

	return ""
}

func getPossibleDigitWords(words []string, char rune, index int) (possibleWords []string) {
	fmt.Println("words:", words)
	fmt.Println("char:", string(char))
	fmt.Println("index:", index)
	for _, word := range words {
		if index < len(word) && string(word[index]) == string(char) {
			possibleWords = append(possibleWords, word)
		}
	}

	return
}

func getLastDigit(line string) string {
	return "0"
}
