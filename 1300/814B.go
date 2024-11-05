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
	a := make([]int, n)
	b := make([]int, n)
	ans := make([]int, n)
	vis := make([]bool, n+1)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	for i := range b {
		fmt.Fscan(in, &b[i])
		if a[i] == b[i] {
			ans[i] = b[i]
			vis[b[i]] = true
		}
	}
	x, y := 0, 0
	for i := 1; i <= n; i++ {
		if !vis[i] {
			if x == 0 {
				x = i
			} else {
				y = i
			}
		}
	}
	id := []int{}
	for i := range ans {
		if ans[i] == 0 {
			id = append(id, i)
		}
	}
	if len(id) == 1 {
		ans[id[0]] = x
	} else {
		if (a[id[0]] == x || a[id[1]] == y) && (b[id[0]] == x || b[id[1]] == y) {
			ans[id[0]], ans[id[1]] = x, y
		} else {
			ans[id[0]], ans[id[1]] = y, x
		}
	}
	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
}
