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

	ans := 0
	for i := range a {
		ans += (i + 1) * (n - i)
		if i < n-1 && a[i] < a[i+1] {
			ans -= (i + 1) * (n - i - 1)
		}
	}
	fmt.Fprintln(out, ans)
}
