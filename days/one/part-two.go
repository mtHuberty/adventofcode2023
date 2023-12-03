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
	file, err := os.Open("./days/one/input.txt")
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
			log.Fatalf("could not convert first digit to int: %s", err)
		}
		if _, err := strconv.Atoi(lastDigitStr); err != nil {
			log.Fatalf("could not convert last digit to int: %s", err)
		}

		concatDigits := firstDigitStr + lastDigitStr
		sum, err := strconv.Atoi(concatDigits)
		if err != nil {
			log.Fatalf("could not convert concat digits to int: %s", err)
		}

		total += sum
	}

	return strconv.Itoa(total)
}

func getFirstDigit(line string) string {
	for i, char := range line {

		// If we find a string representation of a number (1, 2, 3...) we found it
		if _, err := strconv.Atoi(string(char)); err == nil {
			return string(char)
		}

		// Otherwise we have to check for words ("one", "two"...)
		if num := numWordThatStartsAtIndex(line[i:]); num != "" {
			return num
		}
	}

	log.Fatal("hit end of line without finding a digit")

	return ""
}

func getLastDigit(line string) string {
	for i := len(line) - 1; i >= 0; i-- {

		char := line[i]

		// If we find a string representation of a number (1, 2, 3...) we found it
		if _, err := strconv.Atoi(string(char)); err == nil {
			return string(char)
		}

		// Otherwise we have to check for words ("one", "two"...)
		if num := numWordThatStartsAtIndex(line[i:]); num != "" {
			return num
		}
	}

	log.Fatal("hit end of line without finding a digit")

	return ""
}

func numWordThatStartsAtIndex(partialLine string) string {
	possibleDigitWords := make([]string, len(digitWords))
	copy(possibleDigitWords, digitWords)

	for i, char := range partialLine {

		stillValidDigitWords := []string{}

		for _, word := range possibleDigitWords {
			if i < len(word) && string(word[i]) == string(char) {
				if i == len(word)-1 {
					switch possibleDigitWords[0] {
					case "one":
						return "1"
					case "two":
						return "2"
					case "three":
						return "3"
					case "four":
						return "4"
					case "five":
						return "5"
					case "six":
						return "6"
					case "seven":
						return "7"
					case "eight":
						return "8"
					case "nine":
						return "9"
					}
				}

				stillValidDigitWords = append(stillValidDigitWords, word)
			}
		}

		if len(possibleDigitWords) == 0 {
			return ""
		}

		possibleDigitWords = stillValidDigitWords
	}

	return ""
}
