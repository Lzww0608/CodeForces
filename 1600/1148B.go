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

	var n, m, ta, tb, k int
	fmt.Fscan(in, &n, &m, &ta, &tb, &k)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	b := make([]int, m)
	for i := range b {
		fmt.Fscan(in, &b[i])
	}

	if k >= min(n, m) {
		fmt.Fprintln(out, -1)
		return
	}

	ans := 0
	for i, j := 0, 0; i <= k; i++ {
		for j < m && b[j] < a[i]+ta {
			j++
		}
		if k-i >= m-j {
			fmt.Fprintln(out, -1)
			return
		}
		ans = max(ans, b[j+k-i]+tb)
	}
	fmt.Fprintln(out, ans)
}
