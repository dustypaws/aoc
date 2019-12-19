package main

import (
	"bufio"
	"fmt"
	"os"
)

func compare(n, h string) (diff int) {
	for k, v := range n {
		if v != rune(h[k]) {
			diff++
		}
	}
	return
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	var txt []string
	for s.Scan() {
		txt = append(txt, s.Text())
	}
	// Pick next string, starting at the top of the list.
	for _, v := range txt {
		// Compare every other string against the currently chosen one.
		for _, vv := range txt {
			// Make sure we're not comparing the current string against itself.
			if vv != v {
				if compare(v, vv) == 1 {
					fmt.Println(v, vv)
				}
			}
		}
	}
}
