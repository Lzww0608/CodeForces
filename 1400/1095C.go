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

	var n, k int
	fmt.Fscan(in, &n, &k)

	a := make([]int, k)
	for i := range a {
		a[i] = 1
	}
	n -= k
	for i := range a {
		for a[i] <= n {
			n -= a[i]
			a[i] <<= 1
		}

		if n == 0 {
			fmt.Fprintln(out, "YES")
			for _, x := range a {
				fmt.Fprint(out, x, " ")
			}
			return
		}
	}

	fmt.Fprintln(out, "NO")
}
