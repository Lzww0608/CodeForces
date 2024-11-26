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

	var n, k int
	fmt.Fscan(in, &n, &k)
	if k > n*n {
		fmt.Fprintln(out, -1)
		return
	}
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, n)
	}
	t := k
	for i := 0; i < n && t > 0; i++ {
		a[i][i] = 1
		t--
		for j := i + 1; j < n && t > 1; j++ {
			t -= 2
			a[i][j] = 1
			a[j][i] = 1
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fprint(out, a[i][j], " ")
		}
		fmt.Fprintln(out)
	}
}
