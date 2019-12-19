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

var santa_coord = coords{0, 0}
var robo_coord = coords{0, 0}

func main() {

	file, err := os.Open("../input.txt")
	if err != nil {
		logrus.Fatalf("error reading file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)

	var m mapp

	cnt := 0
	for scanner.Scan() {
		cnt++
		turn := "s"
		if cnt%2 == 0 {
			turn = "r"
		}
		m.insertIntoMap(scanner.Text(), turn)
	}
	fmt.Printf("Solution Part II: %d", len(m.m))
}

func (m *mapp) insertIntoMap(direction, turn string) {
	if m.m == nil {
		m.m = make(map[coords]int)
		c := coords{0, 0}
		m.m[c] = 1
	}

	var new_coord coords
	switch direction {
	case UP:
		if turn == "s" {
			new_coord = coords{x: santa_coord.x - 1, y: santa_coord.y}
		} else {
			new_coord = coords{x: robo_coord.x - 1, y: robo_coord.y}
		}
	case DOWN:
		if turn == "s" {
			new_coord = coords{x: santa_coord.x + 1, y: santa_coord.y}
		} else {
			new_coord = coords{x: robo_coord.x + 1, y: robo_coord.y}
		}
	case LEFT:
		if turn == "s" {
			new_coord = coords{x: santa_coord.x, y: santa_coord.y - 1}
		} else {
			new_coord = coords{x: robo_coord.x, y: robo_coord.y - 1}
		}
	case RIGHT:
		if turn == "s" {
			new_coord = coords{x: santa_coord.x, y: santa_coord.y + 1}
		} else {
			new_coord = coords{x: robo_coord.x, y: robo_coord.y + 1}
		}
	}

	m.m[new_coord]++
	if turn == "s" {
		santa_coord = new_coord
	} else {
		robo_coord = new_coord
	}
}
