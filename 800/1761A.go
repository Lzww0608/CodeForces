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

	var t, n, a, b int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n, &a, &b)
		if a+b < n-1 || (a == n && b == n) {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}

	}
}
