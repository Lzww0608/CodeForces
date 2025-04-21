package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 26
const MAXN int = 100_005

var ch [MAXN][N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	cnt := 1

	insert := func(s string) {
		u := 0
		for i := range s {
			x := int(s[i] - 'a')
			if ch[u][x] == 0 {
				ch[u][x] = cnt
				cnt++
			}
			u = ch[u][x]
		}
		return
	}

	var n, k int
	fmt.Fscan(in, &n, &k)
	var s string
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s)
		insert(s)
	}

	win := make([]bool, MAXN)
	lose := make([]bool, MAXN)

	var dfs func(int)
	dfs = func(u int) {
		f := false
		for i := 0; i < N; i++ {
			if ch[u][i] != 0 {
				f = true
				v := ch[u][i]
				dfs(v)
				if !win[v] {
					win[u] = true
				}
				if !lose[v] {
					lose[u] = true
				}
			}
		}
		if !f {
			lose[u] = true
		}
	}
	dfs(0)
	if win[0] && (lose[0] || k&1 == 1) {
		fmt.Fprintln(out, "First")
	} else {
		fmt.Fprintln(out, "Second")
	}
}
