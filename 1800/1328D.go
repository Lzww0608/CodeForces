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
	same, neigh := true, false
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		if a[i] != a[0] {
			same = false
		}
		if i > 0 && a[i] == a[i-1] {
			neigh = true
		}
	}

	if a[0] == a[n-1] {
		neigh = true
	}

	if same {
		fmt.Fprintln(out, 1)
		for range n {
			fmt.Fprintf(out, "%d ", 1)
		}
		fmt.Fprintln(out)
		return
	}

	if n&1 == 0 {
		fmt.Fprintln(out, 2)
		for i := range n {
			fmt.Fprintf(out, "%d ", i&1+1)
		}
		fmt.Fprintln(out)
		return
	} else if neigh {
		fmt.Fprintln(out, 2)
		x := 1
		f := false
		for i := range n {
			if !f && i > 0 && a[i] == a[i-1] {
				f = true
				x ^= 1
			}
			fmt.Fprintf(out, "%d ", x+1)
			x ^= 1

		}
		fmt.Fprintln(out)
		return
	}
	fmt.Fprintln(out, 3)
	fmt.Fprintf(out, "%d ", 3)
	for i := 1; i < n; i++ {
		fmt.Fprintf(out, "%d ", i&1+1)
	}
	fmt.Fprintln(out)
}
