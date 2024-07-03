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
		if n == 3 {
			fmt.Fprintln(out, "NO")
			continue
		}
		fmt.Fprintln(out, "YES")
		if n == 2 {
			fmt.Fprintln(out, 1, -1)
			continue
		}
		for i := 0; i < n; i++ {
			if i&1 == 0 {
				fmt.Fprint(out, (n-2)/2, " ")
			} else {
				fmt.Fprint(out, -((n - 1) / 2), " ")
			}
		}
		fmt.Fprintln(out)

	}
}
