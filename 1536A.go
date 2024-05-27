package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, n int
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		var x, mx int = 0, 0
		f := false
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &x)
			if x < 0 {
				f = true
			} else {
				mx = max(x, mx)
			}

		}
		if f {
			fmt.Fprintln(out, "NO")
			continue
		}
		fmt.Fprintln(out, "YES")
		fmt.Fprintln(out, mx+1)
		for i := 0; i <= mx; i++ {
			fmt.Fprint(out, i, " ")
		}
		fmt.Fprintln(out)
	}
}
