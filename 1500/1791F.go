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

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n, q int
	fmt.Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	f := make([]int, n+1)
	update := func(i, v int) {
		for ; i <= n; i += i & (-i) {
			f[i] += v
		}
	}

	query := func(i int) (res int) {
		for ; i > 0; i -= i & (-i) {
			res += f[i]
		}
		return
	}
	cnt := make([]int, n)

	var x, l, r int
	for range q {
		fmt.Fscan(in, &x)
		if x == 1 {
			fmt.Fscan(in, &l, &r)
			l--
			r--
			update(l+1, 1)
			update(r+2, -1)
		} else {
			fmt.Fscan(in, &l)
			l--
			tmp := query(l + 1)
			cur := tmp - cnt[l]
			cnt[l] = tmp
			for cur > 0 && a[l] >= 10 {
				t := 0
				for a[l] > 0 {
					t += a[l] % 10
					a[l] /= 10
				}
				a[l] = t
				cur--
			}

			fmt.Fprintln(out, a[l])
		}
	}
}
