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
	var n, k int
	fmt.Fscan(in, &n, &k)
	a := make([][2]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i][0])
		a[i][1] = i
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &b[i])
	}
	sort.Ints(b)
	sort.Slice(a, func(i, j int) bool {
		return a[i][0] < a[j][0]
	})
	ans := make([]int, n)
	for i := 0; i < k; i++ {
		if b[i] >= a[n-k+i][0] {
			fmt.Fprintln(out, "NO")
			return
		}
		ans[a[n-k+i][1]] = b[i]
	}
	for i := k; i < n; i++ {
		if b[i] < a[i-k][0] {
			fmt.Fprintln(out, "NO")
			return
		}
		ans[a[i-k][1]] = b[i]
	}
	fmt.Fprintln(out, "YES")
	for i := 0; i < n; i++ {
		fmt.Fprint(out, ans[i], " ")
	}
	fmt.Fprintln(out)
}
