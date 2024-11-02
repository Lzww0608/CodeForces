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
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, n)
	}

	k := 1
	for j := 0; j < n; j++ {
		if j&1 == 0 {
			for i := 0; i < n; i++ {
				a[i][j] = k
				k++
			}
		} else {
			for i := n - 1; i >= 0; i-- {
				a[i][j] = k
				k++
			}
		}

	}
	for i := 0; i < n; i++ {
		for _, x := range a[i] {
			fmt.Fprint(out, x, " ")
		}
		fmt.Fprintln(out)
	}
}
