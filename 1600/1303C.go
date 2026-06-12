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

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

const N int = 26

func solve(in *bufio.Reader, out *bufio.Writer) {
	var s string
	fmt.Fscan(in, &s)
	n := len(s)

	alp := [N*2 + 2]byte{}
	vis := [N]bool{}
	for i := range alp {
		alp[i] = '0'
	}

	var dfs func(int, int, bool) bool
	dfs = func(i, j int, move bool) bool {
		if i == n {
			return true
		}

		if alp[j] == '0' {
			if vis[s[i]-'a'] {
				return false
			}
			vis[s[i]-'a'] = true
			alp[j] = s[i]
			if !dfs(i+1, j, false) {
				alp[j] = '0'
				vis[s[i]-'a'] = false
				return false
			}
			return true
		} else if move {
			if alp[j] != s[i] {
				return false
			}
			return dfs(i+1, j, false)
		} else if !move {
			if dfs(i, j+1, true) {
				return true
			} else if dfs(i, j-1, true) {
				return true
			}
		}

		return false
	}

	if dfs(0, N, false) {
		fmt.Fprintln(out, "YES")
		for i := range alp {
			if alp[i] != '0' {
				fmt.Fprintf(out, "%c", alp[i])
			}
		}
		for i := range vis {
			if !vis[i] {
				fmt.Fprintf(out, "%c", byte(i)+'a')
			}
		}
		fmt.Fprintln(out)
	} else {
		fmt.Fprintln(out, "NO")
	}
}
