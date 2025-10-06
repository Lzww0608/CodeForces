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

	l, r := 0, 0
	for _, x := range a {
		t1, t2 := l+x, r+x
		l = min(t1, t2)
		r = max(abs(t1), abs(t2))
	}

	fmt.Fprintln(out, r)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
