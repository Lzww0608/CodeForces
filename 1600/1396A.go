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

	solve(in, out)
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	if n == 1 {
		fmt.Fprintln(out, 1, 1)
		fmt.Fprintln(out, -a[0])
		fmt.Fprintln(out, 1, 1)
		fmt.Fprintln(out, 0)
		fmt.Fprintln(out, 1, 1)
		fmt.Fprintln(out, 0)
		return
	}

	// a - a * n = -a * (n - 1)
	// a = (y - x) * n - y
	fmt.Fprintln(out, 1, n)
	for i := 0; i < n-1; i++ {
		fmt.Fprint(out, n*(-a[i]), " ")
	}
	fmt.Fprintln(out, 0)

	fmt.Fprintln(out, 1, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fprint(out, (n-1)*a[i], " ")
	}
	fmt.Fprintln(out)

	fmt.Fprintln(out, n, n)
	fmt.Fprintln(out, -a[n-1])
}
