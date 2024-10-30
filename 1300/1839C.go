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
	if a[n-1] == 1 {
		fmt.Fprintln(out, "NO")
		return
	}
	fmt.Fprintln(out, "YES")
	for i := n - 1; i >= 0; i-- {
		if a[i] == 0 && (i == 0 || a[i-1] == 0) {
			fmt.Fprint(out, 0, " ")
		} else {
			r := 0
			i--
			for i >= 0 && a[i] == 1 {
				fmt.Fprint(out, 0, " ")
				i--
				r++
			}
			fmt.Fprint(out, r, " ")
			i++
		}
	}
	fmt.Fprintln(out)
}
