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
	cnt := 0
	for i := 1; i < n; i++ {
		if a[i] != a[i-1] {
			cnt++
		}
	}
	if k == 1 {
		if cnt > 0 {
			fmt.Fprintln(out, -1)
		} else {
			fmt.Fprintln(out, 1)
		}
		return
	}
	fmt.Fprintln(out, max(1, (cnt+k-2)/(k-1)))
}
