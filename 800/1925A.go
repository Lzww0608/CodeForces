package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	//in := bufio.NewScanner(os.Stdin)

	var t, n, k int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &k)
		for i := 0; i < n; i++ {
			for j := 0; j < k; j++ {
				Print(string('a' + j))
			}
		}
		Println()
	}

	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
