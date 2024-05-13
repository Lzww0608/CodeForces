package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t, n, k int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &k)
		if (2*n-2)*2+1 >= k {
			Println((k + 1) / 2)
		} else {
			Println(2 * n)
		}

	}
}
