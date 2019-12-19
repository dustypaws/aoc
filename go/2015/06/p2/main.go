package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

type coord struct {
	x, y int
}
type lights struct {
	m map[coord]int
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var l lights
	for scanner.Scan() {
		curLine := scanner.Text()
		var action, nn1, nn2 string
		var startX, startY, endX, endY int

		if strings.HasPrefix(curLine, "toggle") {
			_, err := fmt.Sscanf(curLine, "%s %d,%d %s %d,%d", &action, &startX, &startY, &nn2, &endX, &endY)
			if err != nil {
				logrus.Fatalf("failed (toggle) with error: %v", err)
			}
		} else {
			_, err := fmt.Sscanf(curLine, "%s %s %d,%d %s %d,%d", &nn1, &action, &startX, &startY, &nn2, &endX, &endY)
			if err != nil {
				logrus.Fatalf("failed (turn) with error: %v", err)
			}
		}
		l.setLights(action, startX, startY, endX, endY)
	}
	fmt.Println("Number of turned on lights: ", l.sumBrightness())
}

func (l *lights) setLights(action string, startX, startY, endX, endY int) {
	if l.m == nil {
		l.m = make(map[coord]int)
	}
	for yy := startY; endY >= yy; yy++ {
		for xx := startX; endX >= xx; xx++ {
			co := coord{x: xx, y: yy}
			switch action {
			case "on":
				l.m[co] += 1
			case "off":
				if l.m[co] > 0 {
					l.m[co]--
				}
			case "toggle":
				l.m[co] += 2
			}
		}
	}
}

func (l *lights) sumBrightness() int {
	brightness := 0
	for _, v := range l.m {
		brightness += v
	}
	return brightness
}
