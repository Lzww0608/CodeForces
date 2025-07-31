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

	ans := make([]int, 0, n)
	ans = append(ans, a[0])
	pre := a[0]
	for i := 1; i < n-1; i++ {
		if a[i] > pre && a[i] < a[i+1] || a[i] < pre && a[i] > a[i+1] {
			continue
		} else {
			ans = append(ans, a[i])
			pre = a[i]
		}
	}
	ans = append(ans, a[n-1])
	fmt.Fprintln(out, len(ans))
	for _, x := range ans {
		fmt.Fprintf(out, "%d ", x)
	}
	fmt.Fprintln(out)
}
