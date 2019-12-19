package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type coords struct {
	x, y int
}
type fabric struct {
	ids []int

	// Map[Coordinates]claims_iterator
	m map[coords]int
}

func main() {
	var c fabric

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	for s.Scan() {
		var id, x, y, w, h int
		_, err := fmt.Sscanf(s.Text(), "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
		if err != nil {
			log.Fatalf("error parsing values from input file: %v", err)
		}
		c.addPatch(id, x, y, w, h)
	}
	c.printMap()
}

func (f *fabric) addPatch(id, x, y, w, h int) {
	if f.m == nil {
		f.m = make(map[coords]int)
	}
	for a := 0; a <= w; a++ {
		for b := 0; b <= h; b++ {
			pos := coords{x + a, y + b}
			f.ids = append(f.ids, id)
			f.m[pos]++
		}
	}
}

func (f *fabric) printMap() {
	for k, id := range f.ids {
		fmt.Printf("[%d] -> %d\n", k, id)
	}
}
