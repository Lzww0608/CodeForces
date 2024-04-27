package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	//in := bufio.NewScanner(os.Stdin)

	var t int
	for Fscan(in, &t); t > 0; t-- {
		var s []byte
		Fscan(in, &s)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		Println(string(s))
	}

	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
