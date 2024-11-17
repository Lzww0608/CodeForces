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

	sum, mx := 0, 0
	ans1, ans2 := 0, 0
	for i := 0; i < n-1; i++ {
		sum += abs(a[i] - a[i+1])
		if i < n-2 {
			ans1 += abs(a[i] - a[i+1])
		}
		if i > 0 {
			ans2 += abs(a[i] - a[i+1])
			mx = max(mx, abs(a[i]-a[i-1])+abs(a[i]-a[i+1])-abs(a[i-1]-a[i+1]))
		}

	}
	fmt.Fprintln(out, min(ans1, ans2, sum-mx))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
