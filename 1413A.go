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
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}

		for i := range a {
			if i&1 == 0 {
				fmt.Fprint(out, a[i+1], " ")
			} else {
				fmt.Fprint(out, -a[i-1], " ")
			}

		}
		fmt.Fprintln(out)
	}
}
