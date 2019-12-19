package main

import (
	"io/ioutil"
	"fmt"
	"regexp"
	"strings"
)

var hasHex = regexp.MustCompile(`\\x..`)
var hasEsc = regexp.MustCompile(`\\"`)
var hasDbl = regexp.MustCompile(`\\\\`)

func main() {
	file, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1: ", memoryDifference(string(file)))
	fmt.Println("Part 2: ", encodingDifference(string(file)))
	// _ = encodingDifference(file)
}

func memoryDifference(s string) int {
	lit := 0
	lines := strings.Split(strings.TrimSpace(s), "\n")
	for _, v := range lines {
		lit += len(v)
		// Order of the following 3 lines is somewhat important.
		l := hasDbl.ReplaceAllString(v, `!`)
		l = hasEsc.ReplaceAllString(l, `#`)
		l = hasHex.ReplaceAllString(l, `$`)
		lit -= len(l) - 2
	}
	return lit
}

func encodingDifference(s string) int {
	lit := 0
	enc := 0
	lines := strings.Split(strings.TrimSpace(s), "\n")
	for _,line:= range lines {
		enc += 2
		lit += len(line)
		for _, v := range line {
			switch v {
			case '"':
				enc += 2
			case '\\':
				enc += 2
			default:
				enc++
			}
		}
	}
	return enc-lit
}
