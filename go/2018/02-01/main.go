package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	counter2, counter3 := 0, 0

	for s.Scan() {
		txt := s.Text()
		seen := make(map[rune]int)
		for _, v := range txt {
			seen[v]++
		}
		had2, had3 := false, false
		for _, vv := range seen {
			if vv == 2 && !had2 {
				counter2++
				had2 = true
			}
			if vv == 3 && !had3 {
				counter3++
				had3 = true
			}
		}
	}
	fmt.Println(counter2 * counter3)
}
