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

		for i, x := range a {
			if i&1 == 0 {
				fmt.Fprint(out, abs(x), " ")
			} else {
				fmt.Fprint(out, -abs(x), " ")
			}

		}
		fmt.Fprintln(out)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
