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

	l, r := 0, 1_000_000_000
	for i := 1; i < n; i++ {
		x, y := a[i-1], a[i]
		if x < y {
			r = min(r, (x+y)/2)
		} else if x > y {
			l = max(l, (x+y+1)/2)
		}
		if l > r {
			fmt.Fprintln(out, -1)
			return
		}
	}
	fmt.Fprintln(out, (l+r)>>1)

}
