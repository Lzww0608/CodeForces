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

	var n int
	fmt.Fscan(in, &n)
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	merge := func(x, y int) bool {
		rx, ry := find(x), find(y)
		if rx != ry {
			fa[rx] = ry
			return true
		}
		return false
	}

	del := [][2]int{}
	for range n - 1 {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		if !merge(u, v) {
			del = append(del, [2]int{u, v})
		}
	}

	add := [][2]int{}
	pre := -1
	for i := range n {
		if i == find(i) {
			if pre != -1 {
				add = append(add, [2]int{pre, i})
			}
			pre = i
		}
	}

	fmt.Fprintln(out, len(del))
	for i := range del {
		fmt.Fprintln(out, del[i][0]+1, del[i][1]+1, add[i][0]+1, add[i][1]+1)
	}
}
