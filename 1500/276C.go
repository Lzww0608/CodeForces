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

	var n, q int
	fmt.Fscan(in, &n, &q)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	f := make([]int, n+1)
	update := func(i, val int) {
		for ; i <= n; i += i & (-i) {
			f[i] += val
		}
	}

	query := func(i int) (res int) {
		for ; i > 0; i -= i & (-i) {
			res += f[i]
		}
		return
	}

	var l, r int
	for ; q > 0; q-- {
		fmt.Fscan(in, &l, &r)
		update(l, 1)
		update(r+1, -1)
	}
	sort.Ints(a)
	ans := 0
	cnt := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		cnt = append(cnt, query(i))
	}
	sort.Ints(cnt)
	for i := n - 1; i >= 0; i-- {
		ans += cnt[i] * a[i]
	}
	fmt.Fprintln(out, ans)
}
