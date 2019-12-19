package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type coords struct {
	x, y int
}

type wireMap struct {
	m map[coords]string
}

var wires [][]string

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(file)

	for s.Scan() {
		wires = append(wires, strings.Split(s.Text(), ","))
	}

	// fmt.Printf("%#v", wires)
	fmt.Println(splitOffOps())
	// fmt.Println(string(wires[0][0][0]))
	// fmt.Println(wires[0][0][1:])
}

func (wm *wireMap) wireMapper() {
	if wm.m == nil {
		wm.m = make(map[coords]string)
	}

}

func getMinMax() (int, int, int) {

	return 0, 0, 0
}

func splitOffOps() ([][]string, [][]int) {
	ops := make([][]string, len(wires)+len(wires[0]))
	steps := make([][]int, len(wires)+len(wires[0]))

	for x, v := range wires {
		for y, i := range v {
			// fmt.Printf("[x:%d, y:%d, O:%v, S:%s]", x, y, string(i[0]), i[1:])
			ops[x][y] = string(i[0])
			steps[x][y], _ = strconv.Atoi(i[1:])
		}
	}
	return ops, steps
}
