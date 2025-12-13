package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, x int32
	fmt.Fscan(in, &n, &m)
	g := make([][]int32, n)
	for i := int32(1); i < n; i++ {
		fmt.Fscan(in, &x)
		x--
		g[x] = append(g[x], int32(i))
	}
	var s string
	fmt.Fscan(in, &s)

	var t int32 = 0
	var inTime []int32 = make([]int32, n)
	var outTime []int32 = make([]int32, n)
	var depTree [][]int32 = make([][]int32, n)
	var dep [][]int32 = make([][]int32, n)
	for i := range dep {
		dep[i] = []int32{0}
	}

	var dfs func(u, fa, d int32)
	dfs = func(u, fa, d int32) {
		x = 1 << int(s[u]-'a')
		t += 1
		inTime[u] = t
		depTree[d] = append(depTree[d], t)
		dep[d] = append(dep[d], dep[d][len(dep[d])-1]^x)
		for _, v := range g[u] {
			if v != fa {
				dfs(v, u, d+1)
			}
		}
		outTime[u] = t
	}
	dfs(0, -1, 0)

	var v, h int
	for range m {
		fmt.Fscan(in, &v, &h)
		v--
		h--
		row := depTree[h]
		l := sort.Search(len(row), func(i int) bool { return row[i] >= inTime[v] })
		r := sort.Search(len(row), func(i int) bool { return row[i] > outTime[v] })
		if x = dep[h][int32(r)] ^ dep[h][int32(l)]; x&(x-1) == 0 {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}
