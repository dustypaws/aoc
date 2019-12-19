package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type dimensions struct {
	l, w, h int
}
type scrap struct {
	s map[int]dimensions
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cnt := 0

	var s scrap

	for scanner.Scan() {
		cnt++
		var l, w, h int
		_, err := fmt.Sscanf(scanner.Text(), "%dx%dx%d", &l, &w, &h)
		if err != nil {
			log.Fatalf("error parsing values from: %v", err)
		}
		s.mapValues(cnt, l, w, h)
		// fmt.Println(cnt, l, w, h)
	}
	fmt.Printf("Part I solution: %d\n", s.calculateTotalSurfaceArea())
	fmt.Printf("Part II solution: %d", s.calculateRibbon())
}

func (s *scrap) mapValues(id, l, w, h int) {
	if s.s == nil {
		s.s = make(map[int]dimensions)
	}
	s.s[id] = dimensions{l, w, h}
}

func findMin(lw, wh, hl int) (min int) {
	set := []int{lw, wh, hl}
	min = set[0]
	for _, v := range set {
		if v < min {
			min = v
		}
	}
	return
}

func (s *scrap) calculateTotalSurfaceArea() (total int) {
	for _, v := range s.s {
		lw := v.l * v.w
		wh := v.w * v.h
		hl := v.h * v.l
		min := findMin(lw, wh, hl)
		total += (2 * lw) + (2 * wh) + (2 * hl) + min
	}
	return
}

func (s *scrap) calculateRibbon() int {
	total := 0
	for _, v := range s.s {
		tmp := []int{v.l, v.w, v.h}
		sort.Ints(tmp)
		total += (tmp[0] + tmp[0] + tmp[1] + tmp[1]) + (tmp[0] * tmp[1] * tmp[2])
	}
	return total
}
