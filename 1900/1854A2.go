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
		var n int
		fmt.Fscan(in, &n)
		a := make([]int, n)

		for i := range a {
			fmt.Fscan(in, &a[i])
		}
		ans := solve(a, n)
		fmt.Fprintln(out, len(ans))
		for _, v := range ans {
			fmt.Fprintln(out, v[0]+1, v[1]+1)
		}
	}
}

func solve(a []int, n int) (ans [][]int) {
	mn, mx := 0, 0
	for i := range a {
		if a[i] > a[mx] {
			mx = i
		}
		if a[i] < a[mn] {
			mn = i
		}
	}

	if a[mx]+a[mn] < 0 {
		for i := 0; i < n; i++ {
			a[i] = -a[i]
		}
		for i := 0; i < n-i-1; i++ {
			a[i], a[n-i-1] = a[n-i-1], a[i]
		}
		ans = solve(a, n)
		for i := range ans {
			ans[i][0], ans[i][1] = n-1-ans[i][0], n-1-ans[i][1]
		}
		return
	}

	cnt := 0
	for i := range a {
		if a[i] < 0 {
			cnt++
		}
	}

	if cnt <= 12 {
		for i := range a {
			if a[i] < 0 {
				ans = append(ans, []int{i, mx})
			}
		}

		for i := 1; i < n; i++ {
			ans = append(ans, []int{i, i - 1})
		}
	} else {
		for i := 0; i < 5; i++ {
			ans = append(ans, []int{mn, mn})
		}
		for i := range a {
			if a[i] > 0 {
				ans = append(ans, []int{i, mn})
			}
		}

		for i := n - 2; i >= 0; i-- {
			ans = append(ans, []int{i, i + 1})
		}
	}

	return
}
