package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func main() {

outer:
	for y := 0; y <= 99999999; y++ {
		key := []byte("bgvyzdsv")
		concat := fmt.Sprintf("%s%d", key, y)
		hash := md5.Sum([]byte(concat))

		fmt.Println(concat, " -> ", fmt.Sprintf("%x", hash))
		if ok := strings.HasPrefix(fmt.Sprintf("%x", hash), "000000"); ok {
			fmt.Println(concat, " -> ", fmt.Sprintf("%x", hash))
			break outer
		}
	}
}
