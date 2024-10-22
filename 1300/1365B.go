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
	var n, x int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	one, zero := 0, 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		if x == 1 {
			one = 1
		} else {
			zero = 1
		}
	}

	f := true
	for i := 1; i < n; i++ {
		if a[i] < a[i-1] {
			f = false
			break
		}
	}

	if f || one == 1 && zero == 1 {
		fmt.Fprintln(out, "Yes")
	} else {
		fmt.Fprintln(out, "No")
	}

}
