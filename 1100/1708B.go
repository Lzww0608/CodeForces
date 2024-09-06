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

	var t, n, l, r int
next:
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n, &l, &r)
		a := make([]int, n+1)
		for i := 1; i <= n; i++ {
			if r/i <= (l-1)/i {
				fmt.Fprintln(out, "NO")
				continue next
			}
			a[i] = r / i * i
		}
		fmt.Fprintln(out, "YES")
		for i := 1; i <= n; i++ {
			fmt.Fprint(out, a[i], " ")
		}
		fmt.Fprintln(out)
	}
}
