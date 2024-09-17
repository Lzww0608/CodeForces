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
		if n > m {
			fmt.Fprintln(out, "NO")
			continue
		}

		if m%n == 0 {
			fmt.Fprintln(out, "YES")
			for i := 0; i < n; i++ {
				fmt.Fprint(out, m/n, " ")
			}
			fmt.Fprintln(out)
		} else if n&1 == 1 {
			fmt.Fprintln(out, "YES")
			fmt.Fprint(out, m-n+1, " ")
			for i := 1; i < n; i++ {
				fmt.Fprint(out, 1, " ")
			}
			fmt.Fprintln(out)
		} else if m&1 == 1 {
			fmt.Fprintln(out, "NO")
		} else {
			fmt.Fprintln(out, "YES")
			fmt.Fprint(out, 1+(m-n)/2, 1+(m-n)/2, " ")
			for i := 2; i < n; i++ {
				fmt.Fprint(out, 1, " ")
			}
			fmt.Fprintln(out)
		}

	}
}
