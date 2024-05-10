package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t, n int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		for i := n; i > 0; i-- {
			fmt.Print(i, " ")
		}
		fmt.Println()
	}
}
