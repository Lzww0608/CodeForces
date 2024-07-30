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
		a := make([]int, n)
		b := make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}

		for i := 0; i < n; {
			j := i
			for j < n && a[i] == a[j] {
				j++
			}
			if j == i+1 {
				fmt.Fprintln(out, -1)
				continue next
			}
			for k := i; k < j; k++ {
				if k == i {
					b[k] = j - 1
				} else {
					b[k] = k - 1
				}
			}
			i = j
		}
		for _, x := range b {
			fmt.Fprint(out, x+1, " ")
		}
		fmt.Fprintln(out)
	}

}
