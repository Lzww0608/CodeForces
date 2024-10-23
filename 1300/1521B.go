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
	pos := 0
	for i := range a {
		fmt.Fscan(in, &a[i])
		if a[i] < a[pos] {
			pos = i
		}
	}
	fmt.Fprintln(out, n-1)
	for i := pos + 1; i < n; i++ {
		fmt.Fprintln(out, pos+1, i+1, a[pos], a[pos]+abs(i-pos))
	}
	for i := pos - 1; i >= 0; i-- {
		fmt.Fprintln(out, pos+1, i+1, a[pos], a[pos]+abs(i-pos))
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
