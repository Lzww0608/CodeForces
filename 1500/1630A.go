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
	var n, k int
	fmt.Fscan(in, &n, &k)
	if k == 0 {
		for i := 0; i < n/2; i++ {
			fmt.Fprintln(out, i, n-1-i)
		}
	} else if k < n-1 {
		fmt.Fprintln(out, k, n-1)
		fmt.Fprintln(out, 0, n-k-1)
		for i := 1; i < n/2; i++ {
			if i != k && i != n-k-1 {
				fmt.Fprintln(out, i, n-i-1)
			}
		}
	} else if n == 4 {
		fmt.Fprintln(out, -1)
	} else {
		fmt.Fprintln(out, n-1, n-2)
		fmt.Fprintln(out, 1, 3)
		fmt.Fprintln(out, 0, n-4)
		for i := 2; i < n/2; i++ {
			if i != 3 {
				fmt.Fprintln(out, i, n-i-1)
			}
		}
	}

}
