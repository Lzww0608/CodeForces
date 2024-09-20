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
		fmt.Fprintln(out, n*3)
		for i, j := 1, n; i < j; i, j = i+1, j-1 {
			for k := 0; k < 3; k++ {
				fmt.Fprintln(out, 1, i, j)
				fmt.Fprintln(out, 2, i, j)
			}
		}

	}

}
