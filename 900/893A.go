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
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	i, j, k := 1, 2, 3
	for _, x := range a {
		if x == k {
			fmt.Fprintln(out, "NO")
			return
		} else if x == i {
			j, k = k, j
		} else if x == j {
			i, k = k, i
		}
	}
	fmt.Fprintln(out, "YES")
}
