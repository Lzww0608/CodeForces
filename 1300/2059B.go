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
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	if k == n {
		for i := 1; i < n; i += 2 {
			if a[i] != (i+1)/2 {
				fmt.Fprintln(out, (i+1)/2)
				return
			}
		}
		fmt.Fprintln(out, n/2+1)
		return
	}

	for i := 1; i < n-(k-2); i++ {
		if a[i] != 1 {
			fmt.Fprintln(out, 1)
			return
		}
	}

	fmt.Fprintln(out, 2)
}
