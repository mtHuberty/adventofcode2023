package util

import (
	"bufio"
	"log"
	"os"
)

type Scanner struct {
	*bufio.Scanner
	file *os.File
}

func (s *Scanner) Close() {
	s.file.Close()
}

func NewScanner(path string) *Scanner {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("could not open file:", err)
	}

	scanner := bufio.NewScanner(file)

	return &Scanner{
		scanner,
		file,
	}
}
