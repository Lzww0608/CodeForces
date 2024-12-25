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

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	sort.Ints(a)
	if a[0] >= 0 {
		fmt.Fprintln(out, "NO")
		return
	}
	d := a[n-1] - a[0]
	ans := make([]int, 0, n)
	cur := 0
	for i, j := n-1, 0; i >= j; {
		if cur+a[i] < d {
			ans = append(ans, a[i])
			cur += a[i]
			i--
		} else {
			if a[j] >= 0 {
				fmt.Fprintln(out, "NO")
				return
			}
			ans = append(ans, a[j])
			cur = max(cur+a[j], 0)
			j++
		}
	}
	fmt.Fprintln(out, "YES")
	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)
}
