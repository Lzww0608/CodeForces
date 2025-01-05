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

	var n int
	fmt.Fscan(in, &n)
	g := make([][]int, n)
	for u := 1; u < n; u++ {
		var v int
		fmt.Fscan(in, &v)
		g[v-1] = append(g[v-1], u)
	}

	a := make([]int, 0, n)
	var f func(int) int
	f = func(v int) int {
		ans := 0
		for _, u := range g[v] {
			ans += f(u)
		}
		if ans == 0 {
			ans = 1
		}
		a = append(a, ans)
		return ans
	}

	f(0)
	sort.Ints(a)
	for _, v := range a {
		fmt.Fprint(out, v, " ")
	}
	fmt.Fprintln(out)
}
