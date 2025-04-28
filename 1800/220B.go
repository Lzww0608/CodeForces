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

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
		if a[i] > n {
			a[i] = n + 1
		}
	}
	q := make([][2]int, m)
	id := make([]int, m)
	for i := range q {
		id[i] = i
		fmt.Fscan(in, &q[i][0], &q[i][1])
		q[i][0]--
		q[i][1]--
	}

	sort.Slice(id, func(i, j int) bool {
		if q[id[i]][0] / 500 == q[id[j]][0] / 500 {
			return q[id[i]][1] < q[id[j]][1]
		}
		return q[id[i]][0] < q[id[j]][0]
	})

	ans := make([]int, m)
	sum, l, r := 0, 0, -1
	cnt := make([]int, n+2)
	for _, i := range id {
		for r < q[i][1] {
			r++
			if cnt[a[r]] == a[r] {
				sum--
			}
			cnt[a[r]]++
			if cnt[a[r]] == a[r] {
				sum++
			}
		}
		for r > q[i][1] {
			if cnt[a[r]] == a[r] {
				sum--
			}
			cnt[a[r]]--
			if cnt[a[r]] == a[r] {
				sum++
			}
			r--
		}

		for l < q[i][0] {
			if cnt[a[l]] == a[l] {
				sum--
			}
			cnt[a[l]]--
			if cnt[a[l]] == a[l] {
				sum++
			}
			l++
		}

		for l > q[i][0] {
			l--
			if cnt[a[l]] == a[l] {
				sum--
			}
			cnt[a[l]]++
			if cnt[a[l]] == a[l] {
				sum++
			}
		}

		//fmt.Fprintln(out, i, l, r, sum)
		ans[i] = sum
	}
	for _, x := range ans {
		fmt.Fprintln(out, x)
	}
}
