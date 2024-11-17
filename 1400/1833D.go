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
	pos := 0
	for pos < n && a[pos] != n {
		pos++
	}

	if pos == 0 {
		for pos < n && a[pos] != n-1 {
			pos++
		}
	}

	x := pos - 1
	if pos == n-1 {
		x++
	}
	for x > 0 && a[x-1] > a[0] {
		x--
	}
	for i := pos; i < n; i++ {
		fmt.Fprint(out, a[i], " ")
	}
	for i := pos - 1; i >= x; i-- {
		fmt.Fprint(out, a[i], " ")
	}
	for i := 0; i < x; i++ {
		fmt.Fprint(out, a[i], " ")
	}
	fmt.Fprintln(out)
}
