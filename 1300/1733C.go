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
	if n == 1 {
		fmt.Fprintln(out, 0)
		return
	}
	x := a[0]
	if (a[0]+a[n-1])&1 == 0 {
		x = a[n-1]
	}
	fmt.Fprintln(out, n-1)
	fmt.Fprintln(out, 1, n)
	for i := 1; i < n-1; i++ {
		if (a[i]+x)&1 == 0 {
			fmt.Fprintln(out, i+1, n)
		} else {
			fmt.Fprintln(out, 1, i+1)
		}
	}

}
