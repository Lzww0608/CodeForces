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
		var s string
		fmt.Fscan(in, &s)
		cnt := 0
		f := false
		for i := 0; i < n; i++ {
			if s[i] == '1' {
				cnt++
				if i+1 < n && s[i+1] == '1' {
					f = true
				}

			}
		}
		if cnt&1 == 1 || (cnt == 2 && f) {
			fmt.Fprintln(out, "NO")
		} else {
			fmt.Fprintln(out, "YES")
		}
	}
}
