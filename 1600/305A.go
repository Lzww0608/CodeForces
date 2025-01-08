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
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	sort.Ints(a)
	ans := []int{}
	vis := make(map[int]bool)
	for _, x := range a {
		if check(x) && !vis[sz(x)] {
			vis[sz(x)] = true
			ans = append(ans, x)
		}
	}

	for _, x := range a {
		if !check(x) && (x < 10 || !vis[sz(x)]) {
			ans = append(ans, x)
			break
		}
	}

	if len(ans) == 0 {
		ans = append(ans, a[0])
	}
	fmt.Fprintln(out, len(ans))
	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
}

func check(x int) bool {
	if x == 0 {
		return true
	}
	for x > 0 {
		if x%10 == 0 {
			return true
		}
		x /= 10
	}
	return false
}

func sz(x int) (res int) {
	for x > 0 {
		res++
		x /= 10
	}
	return
}
