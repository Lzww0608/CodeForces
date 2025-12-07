package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 26

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s []byte
	fmt.Fscan(in, &s)
	n := len(s)

	tree := [26][]int{}
	for i := range N {
		tree[i] = make([]int, n+1)
	}
	update := func(t, i int, val int) {
		for ; i <= n; i += i & (-i) {
			tree[t][i] += val
		}
	}

	query := func(t, i int) (res int) {
		for ; i > 0; i -= i & (-i) {
			res += tree[t][i]
		}
		return
	}

	pre := func(t, l, r int) int {
		return query(t, r) - query(t, l-1)
	}

	for i := range n {
		update(int(s[i]-'a'), i+1, 1)
	}

	var q, t int
	for fmt.Fscan(in, &q); q > 0; q-- {
		fmt.Fscan(in, &t)
		if t == 1 {
			var p int
			var c byte
			fmt.Fscanf(in, "%d %c", &p, &c)

			x := int(s[p-1] - 'a')
			y := int(c - 'a')
			if x == y {
				continue
			}

			update(x, p, -1)
			update(y, p, 1)
			s[p-1] = c
		} else {
			var l, r int
			fmt.Fscan(in, &l, &r)

			cnt := 0
			for i := range N {
				if pre(i, l, r) > 0 {
					cnt++
				}
			}

			fmt.Fprintln(out, cnt)
		}
	}
}
