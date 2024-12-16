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

	var n, x int
	fmt.Fscan(in, &n)
	a := make([][]int, n+1)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		x = n - x
		a[x] = append(a[x], i)
	}

	c := 0
	b := make([]int, n)
	for i := 1; i <= n; i++ {
		sz := len(a[i])
		if sz%i != 0 {
			fmt.Fprintln(out, "Impossible")
			return
		}
		for j := 0; j < sz; j += i {
			for k := 0; k < i; k++ {
				b[a[i][j+k]] = c
			}
			c++
		}
	}
	fmt.Fprintln(out, "Possible")
	for _, x := range b {
		fmt.Fprint(out, x+1, " ")
	}
}
