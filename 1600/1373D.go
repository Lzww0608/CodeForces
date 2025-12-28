package main

import (
	"bufio"
	"fmt"
	"io"
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

func solve(in io.Reader, out io.Writer) {
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	sum := 0
	for i := range n {
		fmt.Fscan(in, &a[i])
		if i&1 == 0 {
			sum += a[i]
		}
	}

	mx, cur := 0, 0
	for i := 1; i < n; i += 2 {
		x := a[i] - a[i-1]
		cur = max(0, cur+x)
		mx = max(mx, cur)
	}
	cur = 0
	for i := 1; i+1 < n; i += 2 {
		x := a[i] - a[i+1]
		cur = max(0, cur+x)
		mx = max(mx, cur)
	}

	fmt.Fprintln(out, sum+mx)
}
