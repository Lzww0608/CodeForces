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
	var n, k int
	fmt.Fscan(in, &n, &k)
	ans := f(n, k)
	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)
}

func f(n, k int) []int {
	if k <= n {
		ans := make([]int, n)
		for i := 0; i < k-1; i++ {
			ans[i] = -1
		}
		if k > 0 {
			ans[k-1] = 200
		}
		if k < n {
			ans[k] = -400
		}
		for i := k + 1; i < n; i++ {
			ans[i] = -1
		}
		return ans
	}
	ans := f(n-1, k-n)
	ans = append(ans, 1000)
	return ans
}
