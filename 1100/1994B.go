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

	var q, n int
next:
	for fmt.Fscan(in, &q); q > 0; q-- {
		fmt.Fscan(in, &n)
		var s, t string
		fmt.Fscan(in, &s, &t)
		f := false
		for i := 0; i < n; i++ {
			if s[i] == '1' {
				f = true
			}
			if t[i] == '1' && !f {
				fmt.Fprintln(out, "NO")
				continue next
			}
		}
		fmt.Fprintln(out, "YES")
	}
}
