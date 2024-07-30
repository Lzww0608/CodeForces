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
		k := 0
		for 1<<(k+1) <= n-1 {
			k++
		}

		for i := (1 << k) - 1; i >= 0; i-- {
			fmt.Fprint(out, i, " ")
		}
		for i := 1 << k; i < n; i++ {
			fmt.Fprint(out, i, " ")
		}
		fmt.Fprintln(out)
	}
}
