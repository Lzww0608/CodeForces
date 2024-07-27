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
		fmt.Fprintln(out, 2)
		fmt.Fprintln(out, n, " ", n-1)
		for i := n; i >= 3; i-- {
			fmt.Fprintln(out, i, " ", i-2)
		}
	}
}
