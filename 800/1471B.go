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
		if n == 3 {
			fmt.Println(-1)
			continue
		}
		fmt.Print(n, n-1, " ")
		for i := 1; i <= n-2; i++ {
			fmt.Print(i, " ")
		}
		fmt.Println()
	}
}
