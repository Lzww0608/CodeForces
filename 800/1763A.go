package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 10

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t, n, x int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)

		mx, mn := 0, 1<<N-1
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &x)
			mx |= x
			mn &= x
		}

		fmt.Fprintln(out, mx-mn)
	}
}
