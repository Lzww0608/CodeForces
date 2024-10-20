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

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	sum := 0
	for i := range a {
		fmt.Fscan(in, &a[i])
		sum += a[i]
	}
	f := true
	for sum > 0 {
		for i := 0; sum > 0 && i < n-1; {
			if a[i] > 0 && f {
				a[i]--
				sum--
				f = false
				fmt.Fprint(out, "P")
			} else {
				f = true
				fmt.Fprint(out, "R")
				i++
			}
		}
		for i := n - 1; sum > 0 && i > 0; {
			if a[i] > 0 && f {
				a[i]--
				sum--
				f = false
				fmt.Fprint(out, "P")
			} else {
				f = true
				fmt.Fprint(out, "L")
				i--
			}
		}
	}
	fmt.Fprintln(out)
}
