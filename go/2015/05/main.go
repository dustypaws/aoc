package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	niceStrings := 0
	cnt := 0
	for scanner.Scan() {
		cnt++
		curLine := scanner.Text()
		if hasVowels(curLine) &&
			foundLetterTwiceInARow(curLine) &&
			!hasBadStrings(curLine) {
			niceStrings++
			// fmt.Println(niceStrings, "\t(", cnt, ")", ": ", curLine)
		}
	}
	fmt.Printf("Number of nice strings: %d", niceStrings)
}

func hasVowels(s string) bool {
	vowels := []string{"a", "e", "i", "o", "u"}
	foundNumOfVowels := 0
	for _, v := range vowels {
		foundNumOfVowels += strings.Count(s, v)
	}
	return foundNumOfVowels >= 3
}

func foundLetterTwiceInARow(s string) bool {
	var prevLetter string
	foundDoubles := 0
	for _, v := range s {
		if prevLetter == string(v) {
			foundDoubles++
		}
		prevLetter = string(v)

	}
	return foundDoubles >= 1
}

func hasBadStrings(s string) bool {
	badStrings := []string{"ab", "cd", "pq", "xy"}
	for _, v := range badStrings {
		if strings.Contains(s, v) {
			return true
		}
	}
	return false
}
