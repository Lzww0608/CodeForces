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
		var s, t string
		fmt.Fscan(in, &s)
		fmt.Fscan(in, &t)
		f := make([]int, n+1)
		for i := 1; i <= n; i++ {
			has0, has1 := false, false
			for j := i - 1; j >= 0 && j >= i-3; j-- {
				if s[j] == '0' || t[j] == '0' {
					has0 = true
				}
				if s[j] == '1' || t[j] == '1' {
					has1 = true
				}

				mex := 0
				if has0 {
					mex = 1
				}
				if has1 && has0 {
					mex = 2
				}
				f[i] = max(f[i], f[j]+mex)
			}
		}

		fmt.Fprintln(out, f[n])
	}
}
