package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

var masses []int

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	fmt.Println("Part 1: ", part1(s))
	fmt.Println("Part 2: ", part2(masses))
}

func part1(s *bufio.Scanner) int {
	acc := 0
	for s.Scan() {
		ln, _ := strconv.Atoi(s.Text())
		// Need to write all values into an array at the 1st pass,
		// since there seems to be no way to reset the friggin pointer...WTF!
		masses = append(masses, ln)
		acc += calc(ln)
	}
	return acc
}

func part2(s []int) int {
	accumulatedFuelConsumption := 0
	for _, i := range s {
		// Calculate the initial consumption for each module.
		initialFuelConsumption := calc(i)
		// add the initial consumption to the accumulator.
		accumulatedFuelConsumption += initialFuelConsumption
		// initialize the start value for the recursive loop.
		currentStep := initialFuelConsumption

		// calculate the consumption for the fuel recursively.
		for (currentStep/3)-2 >= 0 {
			tmp := calc(currentStep)
			accumulatedFuelConsumption += tmp
			// set currentStep to the just calculated new value and start over.
			currentStep = tmp
		}
	}
	return accumulatedFuelConsumption
}

func calc(i int) int {
	return (int(math.Trunc(float64(i))) / 3) - 2
}
