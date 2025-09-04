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

	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	t, l := 0, 0
	for i := 1; i < k-1; i++ {
		if a[i] > a[i-1] && a[i] > a[i+1] {
			t++
		}
	}

	cur := t
	for i := 1; i+k-1 < n; i++ {
		if a[i] > a[i-1] && a[i] > a[i+1] {
			cur--
		}
		if a[i+k-2] > a[i+k-3] && a[i+k-2] > a[i+k-1] {
			cur++
		}

		if cur > t {
			t, l = cur, i
		}
	}

	fmt.Fprintln(out, t+1, l+1)
}
