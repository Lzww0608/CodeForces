package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	//in := bufio.NewScanner(os.Stdin)

	var t, n int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		x := 1
		for i := 0; i < n; i++ {
			Print(x, " ")
			x += 2
		}
		Println()
	}

	return
}
