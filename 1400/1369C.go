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
	w := make([]int, k)
	for i := range w {
		fmt.Fscan(in, &w[i])
	}
	sort.Ints(w)

	sort.Ints(a)
	ans := 0

	for i := n - k; i < n; i++ {
		ans += a[i]
	}

	for i, j := 0, n-k; i < k; i++ {
		if w[i] == 1 {
			ans += a[n-i-1]
		} else {
			j -= w[i] - 1
			ans += a[j]
		}
	}

	fmt.Fprintln(out, ans)
}
