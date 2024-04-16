package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	lines, err := readLines("sample.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	var userStructSlice []string
	for _, line := range lines {
		userStructSlice = append(userStructSlice, line)
	}

	for i, userStruct := range userStructSlice {
		fmt.Printf("User %d: %s\n", i+1, userStruct)
	}
}
