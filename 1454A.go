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
		for i := 2; i <= n; i++ {
			Print(i, " ")
		}
		Print(1)
		Println()
	}

	return
}
