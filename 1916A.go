package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t, n, k, x int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &k)
		div := 1
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			div *= x
		}
		if 2023%div != 0 {
			Println("NO")
			continue
		}
		Println("YES")
		Print(2023/div, " ")
		for i := 1; i < k; i++ {
			Print(1, " ")
		}
		Println()
	}
}
