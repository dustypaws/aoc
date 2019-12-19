package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		tmp := scanner.Text()
		n, _ := strconv.Atoi(tmp[1:])
		switch tmp[:1] {
		case "0":
			break
		case "+":
			i += n
			break
		case "-":
			i -= n
			break
		}
	}
	fmt.Println(i)
}
