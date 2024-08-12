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

	var t, n, k int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n, &k)
		if k == 1 {
			fmt.Fprintln(out, "YES")
			for i := 0; i < n; i++ {
				fmt.Fprintln(out, i+1)
			}
			continue
		}
		if n&1 == 1 {
			fmt.Fprintln(out, "NO")
			continue
		}
		fmt.Fprintln(out, "YES")
		for i := 1; i <= n; i++ {
			for j := 0; j < k; j++ {
				fmt.Fprint(out, i+n*j, " ")
			}
			fmt.Fprintln(out)
		}
	}
}
