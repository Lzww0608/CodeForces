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

	type pair struct{ x, y int }
	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	a := make([]int, n)
	id := make([]int, n)
	pos := make([]pair, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		a[i]--
		id[a[i]] = i
		pos[i] = pair{i / k, i % k}
	}

	b := make([]int, m)
	ans := 0
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &b[i])
		b[i]--
		ans += pos[id[b[i]]].x + 1
		if id[b[i]] == 0 {
			continue
		}
		x, y := a[id[b[i]]-1], b[i]
		if pos[id[b[i]]].x > 0 || pos[id[b[i]]].y > 0 {
			id[x], id[y] = id[y], id[x]
			a[id[x]], a[id[y]] = a[id[y]], a[id[x]]
			//pos[id[b[i]]], pos[id[a[id[b[i]]-1]]] = pos[id[a[id[b[i]]-1]]], pos[id[b[i]]]
		}
		//fmt.Fprintln(out, ans)
		//fmt.Fprintln(out, id)
	}

	fmt.Fprintln(out, ans)
}
