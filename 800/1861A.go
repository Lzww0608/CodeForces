package main

import (
	"fmt"
	"strings"
)

func main() {
	var t int
	var s string
	for fmt.Scan(&t); t > 0; t-- {
		fmt.Scan(&s)
		if strings.Index(s, "1") < strings.Index(s, "3") {
			fmt.Println(13)
		} else {
			fmt.Println(31)
		}
	}
}
