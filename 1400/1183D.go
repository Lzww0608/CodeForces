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
	var n, x int
	fmt.Fscan(in, &n)
	a := make([]int, n+1)
	cnt := make([]int, n+1)
	for range n {
		fmt.Fscan(in, &x)
		a[x]++
	}

	for _, y := range a {
		cnt[y]++
	}

	ans := 0
	x = 0
	for i := n; i > 0; i-- {
		x += cnt[i]
		if x > 0 {
			x--
			ans += i
		}
	}

	fmt.Fprintln(out, ans)
}
