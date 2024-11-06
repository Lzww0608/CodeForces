package main

import (
	"bufio"
	"fmt"
	"math"
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
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	if n&1 == 0 {
		ans := 0
		for i := 1; i < n; i += 2 {
			ans = max(ans, a[i]-a[i-1])
		}
		fmt.Fprintln(out, ans)
		return
	}

	ans := math.MaxInt64
	for i := 0; i < n; i++ {
		cur, pre := 1, -1
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if pre == -1 {
				pre = a[j]
			} else {
				cur = max(cur, a[j]-pre)
				pre = -1
			}
		}
		ans = min(ans, cur)
	}
	fmt.Fprintln(out, ans)
}
