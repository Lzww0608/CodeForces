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

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	ans := 0
	for i := 0; i < n-1; i++ {
		a[i] = abs(a[i] - a[i+1])
	}

	f := func(i int) {
		cur := 0
		for j := i; j < n-1; j++ {
			x := a[j]
			if j&1 != i&1 {
				x = -x
			}
			cur = max(cur+x, x)
			ans = max(ans, cur)
		}
	}
	f(0)
	f(1)
	fmt.Fprintln(out, ans)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
