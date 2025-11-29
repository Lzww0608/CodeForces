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

	solve(in, out)

}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	sum := 0
	for i := range a {
		fmt.Fscan(in, &a[i])
		sum += a[i]
	}
	sort.Ints(a)

	var m, x, y int
	for fmt.Fscan(in, &m); m > 0; m-- {
		fmt.Fscan(in, &x, &y)
		p := sort.SearchInts(a, x)
		ans := 0
		if p == n {
			ans += x - a[n-1]
			ans += max(0, y-(sum-a[n-1]))
		} else {
			ans = max(0, y-(sum-a[p]))
			if p > 0 {
				ans = min(ans, max(0, y-(sum-a[p-1]))+x-a[p-1])
			}
		}

		fmt.Fprintln(out, ans)
	}
}
