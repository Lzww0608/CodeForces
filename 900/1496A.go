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
		var s string
		fmt.Fscan(in, &s)
		ok, i := true, 0
		for i = 0; i < k; i++ {
			ok = ok && (s[i] == s[n-i-1])
		}
		if ok && k*2 < n {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
