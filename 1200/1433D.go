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
		a := make([]int, n)
		f := false
		idx := 0
		for i := range a {
			fmt.Fscan(in, &a[i])
			if a[i] != a[0] {
				idx = i
				f = true
			}
		}
		if !f {
			fmt.Fprintln(out, "NO")
			continue
		}
		fmt.Fprintln(out, "YES")
		m := make([]map[int]int, n)
		for i := range m {
			m[i] = make(map[int]int)
		}
		vis := make([]bool, n)
		var dfs func(idx, pre int)
		dfs = func(idx, pre int) {
			if idx >= n || vis[idx] {
				return
			}
			if pre != -1 {
				fmt.Fprintln(out, pre+1, idx+1)
			}
			vis[idx] = true
			for i := 0; i < n; i++ {
				if !vis[i] && a[i] != a[idx] {
					dfs(i, idx)
				}
			}

		}
		dfs(idx, -1)
	}
}
