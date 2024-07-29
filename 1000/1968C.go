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
		a := make([]int, n-1)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}
		x := int(1e8)
		fmt.Fprint(out, x, " ")
		for _, t := range a {
			x += t
			fmt.Fprint(out, x, " ")
		}
		fmt.Fprintln(out)

	}

}
