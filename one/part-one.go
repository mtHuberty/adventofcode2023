package one

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

// The code reads input from a file (input.txt) that contains a series of strings, each representing a number.
// For each line, it finds the first and last digit and calculates their sum.
// The sum of these pairs is accumulated to calculate the final result.
// It uses goroutines and a sync.WaitGroup to concurrently find the first and last digits in each string.
// The reverse function is used to reverse a string, and it's employed to find the last digit efficiently.
// The solution uses error handling to log and exit if any unexpected conditions occur.
//
// Overall, the code efficiently processes the input strings,
// finds pairs of digits, calculates their sum, and accumulates
// the total sum for all pairs. The goroutines enhance concurrency
// for processing multiple strings simultaneously.

func SolvePartOne() string {
	file, err := os.Open("one/input.txt")
	if err != nil {
		log.Fatal(fmt.Errorf("could not open file: %w", err))
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		var firstDigit *int
		var lastDigit *int

		wg := sync.WaitGroup{}
		wg.Add(2)
		go func() {
			defer wg.Done()
			for _, char := range line {
				if digit, err := strconv.Atoi(string(char)); err == nil {
					firstDigit = &digit
					return
				}
			}

			fmt.Println("no digit found in string:", line)
			os.Exit(1)
		}()
		go func() {
			defer wg.Done()
			for _, char := range reverse(line) {
				if digit, err := strconv.Atoi(string(char)); err == nil {
					lastDigit = &digit
					return
				}
			}

			fmt.Println("no digit found in string:", line)
			os.Exit(1)
		}()
		wg.Wait()

		if firstDigit == nil || lastDigit == nil {
			fmt.Println("missing digit, add more logs and retry")
		}

		concatDigits := strconv.Itoa(*firstDigit) + strconv.Itoa(*lastDigit)
		sum, err := strconv.Atoi(concatDigits)
		if err != nil {
			fmt.Println(fmt.Errorf("error converting sum string ot int: %w", err))
			os.Exit(1)
		}

		total += sum
	}

	return strconv.Itoa(total)
}

// Reversing the line isn't the most efficient (2n),
// could just do a regular for loop and go backwareds,
// but that's a bit less readable/consistent.
func reverse(str string) (reversed string) {
	for _, v := range str {
		reversed = string(v) + reversed
	}

	return reversed
}
