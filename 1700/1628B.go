package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	vis := make(map[string]bool)
	str := make([]string, n)
	for i := range str {
		fmt.Fscan(in, &str[i])
	}
	for _, s := range str {
		m := len(s)
		if m == 1 {
			fmt.Fprintln(out, "YES")
			return
		} else if m == 2 {
			if s[0] == s[1] {
				fmt.Fprintln(out, "YES")
				return
			}

			t := string(s[1]) + string(s[0])
			if vis[t] {
				fmt.Fprintln(out, "YES")
				return
			}

			for c := 'a'; c <= 'z'; c++ {
				if vis[t+string(c)] {
					fmt.Fprintln(out, "YES")
					return
				}
			}
		} else {
			t := string(s[2]) + string(s[1])
			tt := []byte(s)
			slices.Reverse(tt)
			rev := string(tt)
			if s[0] == s[2] || vis[t] || vis[rev] {
				fmt.Fprintln(out, "YES")
				return
			}

		}
		vis[s] = true
	}

	fmt.Fprintln(out, "NO")
}
