package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var nums []int
	for scanner.Scan() {
		var n int
		_, err := fmt.Sscanf(scanner.Text(), "%d", &n)
		if err != nil {
			log.Fatalf("could not read %s: %v", scanner.Text(), err)
		}
		nums = append(nums, n)
	}

	sum := 0
	seen := map[int]bool{0: true}
	for {
		for _, n := range nums {
			sum += n
			if seen[sum] {
				fmt.Println(sum)
				return
			}
			seen[sum] = true
		}
	}
}
