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
	ans := 0
	for i := range a {
		fmt.Fscan(in, &a[i])
		ans += a[i] / k
		a[i] %= k
	}

	sort.Ints(a)
	for i, j := 0, n-1; i < j; i++ {
		if a[i]+a[j] >= k {
			ans++
			j--
		}
	}

	fmt.Fprintln(out, ans)
}
