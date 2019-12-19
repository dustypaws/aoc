package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var floor int
	firstHitOnBasement := false

	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tmp := scanner.Text()
		for k, v := range tmp {
			switch v {
			case rune('('):
				floor++
			case rune(')'):
				floor--
			}

			// Part II
			if floor == -1 && !firstHitOnBasement {
				firstHitOnBasement = true
				fmt.Println(k + 1)
			}
		}
	}
	fmt.Println(floor)
}
