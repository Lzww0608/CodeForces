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
next:
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n, &k)
		for i := 0; i <= n; i++ {
			t := (n-1-i)*(n-i)/2 + (i-1)*i/2
			if t == k {
				fmt.Fprintln(out, "YES")
				for j := 0; j < n; j++ {
					if j < i {
						fmt.Fprint(out, -1, " ")
					} else {
						fmt.Fprint(out, 1, " ")
					}
				}
				fmt.Fprintln(out)
				continue next
			}
		}
		fmt.Fprintln(out, "NO")
	}
}
