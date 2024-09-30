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
	var n, d, h int
	fmt.Fscan(in, &n, &d, &h)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	ans := float64(d*h*n) / 2
	for i := 0; i < n-1; i++ {
		if a[i]+h <= a[i+1] {
			continue
		}
		t := float64(h - a[i+1] + a[i])
		ans -= t * t / float64(h*h) * float64(h*d) / 2
	}
	fmt.Fprintln(out, ans)
}
