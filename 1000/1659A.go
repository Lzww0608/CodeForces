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

	var t, n, b, r int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n, &r, &b)
		t := r / (b + 1)
		q := r % (b + 1)
		for i := 0; i < b; i++ {
			for j := 0; j < t; j++ {
				r--
				fmt.Fprint(out, "R")
			}
			if q > 0 {
				q--
				r--
				fmt.Fprint(out, "R")
			}
			fmt.Fprint(out, "B")
		}
		for r > 0 {
			r--
			fmt.Fprint(out, "R")
		}
		fmt.Fprintln(out)
	}
}
