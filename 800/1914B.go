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
		x := n
		for i := 0; i < n; i++ {
			if k > 0 {
				Print(i+1, " ")
				k--
			} else {
				Print(x, " ")
				x--
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
