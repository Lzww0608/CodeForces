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
	sum := 0
	for i := range a {
		fmt.Fscan(in, &a[i])
		sum += a[i]
	}
	cur := 0
	for i := 1; i < n; i += 2 {
		cur += a[i] - 1
	}
	if cur*2 <= sum {
		for i := 0; i < n; i++ {
			if i&1 == 0 {
				fmt.Fprint(out, a[i], " ")
			} else {
				fmt.Fprint(out, 1, " ")
			}
		}
		fmt.Fprintln(out)
	} else {
		for i := 0; i < n; i++ {
			if i&1 == 1 {
				fmt.Fprint(out, a[i], " ")
			} else {
				fmt.Fprint(out, 1, " ")
			}
		}
		fmt.Fprintln(out)
	}
}
