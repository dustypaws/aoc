package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type coords struct {
	x, y int
}
type mapp struct {
	m map[coords]int
}

const (
	UP    = "^"
	DOWN  = "v"
	LEFT  = "<"
	RIGHT = ">"
)

var current_coord = coords{0, 0}

func main() {

	file, err := os.Open("./input.txt")
	if err != nil {
		logrus.Fatalf("error reading file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)

	var m mapp

	for scanner.Scan() {
		m.insertIntoMap(scanner.Text())
	}
	fmt.Printf("Solution Part I: %d", len(m.m))
}

func (m *mapp) insertIntoMap(direction string) {
	if m.m == nil {
		m.m = make(map[coords]int)
		c := coords{0, 0}
		m.m[c] = 1
	}

	var new_coord coords
	switch direction {
	case UP:
		new_coord = coords{x: current_coord.x - 1, y: current_coord.y}
	case DOWN:
		new_coord = coords{x: current_coord.x + 1, y: current_coord.y}
	case LEFT:
		new_coord = coords{x: current_coord.x, y: current_coord.y - 1}
	case RIGHT:
		new_coord = coords{x: current_coord.x, y: current_coord.y + 1}
	}

	m.m[new_coord]++
	current_coord = new_coord
}
