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

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		a[i] += a[i-1]
	}
	ans := 0
	var l, r int
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &l, &r)
		ans += max(0, a[r]-a[l-1])
	}
	fmt.Fprintln(out, ans)
}
