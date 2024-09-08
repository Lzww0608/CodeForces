package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t, n int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		a, b := 1, n-1
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				a, b = n/i, n-n/i
				break
			}
		}
		fmt.Fprintln(out, a, b)
	}
}
