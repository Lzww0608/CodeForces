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
	ans := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	for i, j := 0, 0; i < n; i++ {
		for a[j] < (i - j + 1) {
			j++
		}

		ans[i] = i - j + 1
	}

	for _, x := range ans {
		fmt.Fprintf(out, "%d ", x)
	}
	fmt.Fprintln(out)
}
