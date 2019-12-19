package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	f, _ := ioutil.ReadAll(file)

	ops := convertToInt(strings.Split(string(f), ","))
	// fmt.Println("Part 1:", part1(ops))
	fmt.Print(part2(ops))
}

func convertToInt(s []string) []int {
	var out []int
	for _, i := range s {
		tmp, _ := strconv.Atoi(i)
		out = append(out, tmp)
	}
	return out
}

func part1(ops []int) int {
	lastIter := false

outer:
	for i := 0; i <= len(ops); i += 4 {

		if lastIter || ops[i] > len(ops) {
			break outer
		}

		switch ops[i] {
		case 1:
			ops[ops[i+3]] = add(ops[ops[i+1]], ops[ops[i+2]])
			break
		case 2:
			ops[ops[i+3]] = multi(ops[ops[i+1]], ops[ops[i+2]])
			break
		case 99:
			lastIter = true
		default:
			fmt.Println("How did we get here?", ops[i], "( at Index:", i, ")")
		}
	}
	ops = ops
	return ops[0]
}

func part2(ops []int) (int, int) {
	var i, j int
loop:
	for i = 0; i <= 99; i++ {
		for j = 0; j <= 99; j++ {
			opsCpy := append(ops[:0:0], ops...)
			opsCpy[1], opsCpy[2] = i, j
			nut := part1(opsCpy)
			if nut == 19690720 {
				break loop
			}
		}
	}
	return i, j
}

func add(a, b int) int {
	return a + b
}

func multi(a, b int) int {
	return a * b
}
