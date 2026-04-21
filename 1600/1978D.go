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
	var n, c int
	fmt.Fscan(in, &n, &c)
	a := make([]int, n)
	mx := 0
	for i := range a {
		fmt.Fscan(in, &a[i])
		if a[i] > mx {
			mx = a[i]
		}
	}

	if n == 1 {
		fmt.Fprintln(out, 0)
		return
	}

	sum := c
	pre := a[0] + c
	f := true
	for i, x := range a {
		if x == mx && pre < mx && f {
			fmt.Fprintf(out, "%d ", 0)
		} else {
			sum += x
			if sum >= mx {
				fmt.Fprintf(out, "%d ", i)
			} else {
				fmt.Fprintf(out, "%d ", i+1)
			}
		}

		if x == mx {
			f = false
		}
	}

	fmt.Fprintln(out)
}
