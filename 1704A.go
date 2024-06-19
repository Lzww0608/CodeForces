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

	var t, n, m int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n, &m)
		var a, b string
		fmt.Fscan(in, &a)
		fmt.Fscan(in, &b)

		i, j := n-1, m-1
		for j >= 0 && i >= 0 {
			if a[i] == b[j] {
				j--
			} else if j > 0 {
				break
			}
			i--
		}
		if j < 0 {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}

	}
}
