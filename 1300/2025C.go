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
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	sort.Ints(a)
	ans := 0
	for i, j := 0, 0; j < n && i < n; i++ {
		for j < n && a[j] < a[i]+k && (j <= i || a[j] == a[j-1] || a[j] == a[j-1]+1) {
			j++
		}
		ans = max(ans, j-i)
	}

	fmt.Fprintln(out, ans)
}
