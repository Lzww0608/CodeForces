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

	var t, n int
next:
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		d := make([]int, n)
		for i := range d {
			fmt.Fscan(in, &d[i])
		}
		a := make([]int, n)
		a[0] = d[0]
		for i := 1; i < n; i++ {
			if d[i] != 0 && a[i-1]-d[i] >= 0 {
				fmt.Fprintln(out, -1)
				continue next
			} else {
				a[i] = a[i-1] + d[i]
			}
		}
		for _, x := range a {
			fmt.Fprint(out, x, " ")
		}
		fmt.Fprintln(out)
	}

}
