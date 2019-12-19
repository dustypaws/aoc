package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	cnt, hits := 0, 0
	for scanner.Scan() {
		cnt++
		curLine := scanner.Text()

		if findDoubles(curLine) && findSeperatedByOne(curLine) {
			hits++
			fmt.Println(" ->", curLine)
		}
	}
	fmt.Println("Number of hits: ", hits)
}

func findDoubles(s string) bool {
	prevLetter := ""
	cnt := 0
	for i := 0; i < len(s); i++ {
		if prevLetter == "" || i < 1 {
			prevLetter = string(s[i])
			continue
		}
		if strings.Count(s, fmt.Sprintf("%c%c", s[i-1], s[i])) > 1 {
			// fmt.Println("-> ", fmt.Sprintf("%c%c", s[i-1], s[i]))
			cnt++
		}
	}
	return cnt > 0
}

func findSeperatedByOne(s string) bool {
	cnt := 0
	for i := 1; i < len(s)-1; i++ {
		if s[i-1] == s[i+1] {
			cnt++
			// fmt.Print(cnt, ": sep")
		}
	}
	return cnt >= 1
}
