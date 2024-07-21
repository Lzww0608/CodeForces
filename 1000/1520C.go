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
		if n == 2 {
			fmt.Fprintln(out, -1)
			continue
		}

		for i := 0; i < n; i++ {
			if i&1 == 0 {
				for j := 0; j < n; j++ {
					fmt.Fprint(out, i+1+j*n, " ")
				}
			} else {
				fmt.Fprint(out, n*n-(n-i-1), " ")
				for j := 0; j < n-1; j++ {
					fmt.Fprint(out, i+1+j*n, " ")
				}
			}
			fmt.Fprintln(out)
		}
	}
}
