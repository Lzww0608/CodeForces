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
	pos := make([]int, n+1)
	for i := range a {
		fmt.Fscan(in, &a[i])
		pos[a[i]] = i
	}

	ans := make([]int, n)
	for i := range ans {
		ans[i] = n + 1 - a[i]
	}

	for i := 1; i <= n; i++ {
		if pos[i]&1 != pos[1]&1 {
			ans[pos[i]], ans[pos[1]] = ans[pos[1]], ans[pos[i]]
		}
	}
	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)
}
