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
	var a, b, q int
	fmt.Fscan(in, &a, &b, &q)

	f := make([]int, a*b+1)
	for i := 1; i <= a*b; i++ {
		if i%a%b != i%b%a {
			f[i] = f[i-1] + 1
		} else {
			f[i] = f[i-1]
		}
	}

	x := a * b
	calc := func(n int) int {
		return n/x*f[x-1] + f[n%x]
	}

	var l, r int
	for range q {
		fmt.Fscan(in, &l, &r)
		fmt.Fprintf(out, "%d ", calc(r)-calc(l-1))
	}
	fmt.Fprintln(out)
}
