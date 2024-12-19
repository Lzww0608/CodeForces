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
	var n, q int
	fmt.Fscan(in, &n, &q)
	a := make([]int, n)
	ans := make([]byte, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
		ans[i] = '0'
	}

	cur := 0
	for i := n - 1; i >= 0; i-- {
		if a[i] <= cur {
			ans[i] = '1'
		} else if cur < q {
			ans[i] = '1'
			cur++
		}
	}

	fmt.Fprintln(out, string(ans))
}
