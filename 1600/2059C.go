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
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, n)
		for j := n - 1; j >= 0; j-- {
			fmt.Fscan(in, &a[i][j])
		}
		pre := 0
		for j := range n {
			x := a[i][j]
			a[i][j] = pre
			pre += x
		}
	}

	ans := n
	for j := range n {
		cnt := 0
		for i := range n {
			if a[i][j] == j {
				cnt++
			}
		}

		ans = min(ans, cnt+j)
	}

	fmt.Fprintln(out, ans)
}
