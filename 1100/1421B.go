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
		s := make([]string, n)
		for i := range s {
			fmt.Fscan(in, &s[i])
		}
		ans := 0
		if s[0][1] == s[1][0] && s[n-1][n-2] == s[n-2][n-1] && s[0][1] == s[n-1][n-2] {
			fmt.Fprintln(out, 2)
			fmt.Fprintln(out, 1, 2)
			fmt.Fprintln(out, 2, 1)
			continue
		}
		if s[0][1] != s[1][0] {
			ans++
		}
		if s[n-1][n-2] != s[n-2][n-1] {
			ans++
		}
		fmt.Fprintln(out, ans)
		if s[0][1] != s[1][0] {
			if ans == 1 {
				if s[0][1] != s[n-1][n-2] {
					fmt.Fprintln(out, 2, 1)
				} else {
					fmt.Fprintln(out, 1, 2)
				}
			} else if ans == 2 {
				fmt.Fprintln(out, 2, 1)
			}
		}
		if s[n-1][n-2] != s[n-2][n-1] {
			if s[0][1] != s[n-1][n-2] {
				fmt.Fprintln(out, n-1, n)
			} else {
				fmt.Fprintln(out, n, n-1)
			}
		}

	}
}
