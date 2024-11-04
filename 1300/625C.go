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
	a := make([][]int, n)
	ans := 0
	for j := 0; j < n; j++ {
		ans += n*n - (n - k) - j*(n-k+1)
	}
	id := n * n
	for i := range a {
		a[i] = make([]int, n)
		for j := n - 1; j >= k-1; j-- {
			a[i][j] = id
			id--
		}
	}
	id = 1
	for i := 0; i < n; i++ {
		for j := 0; j < k-1; j++ {
			a[i][j] = id
			id++
		}
	}
	fmt.Fprintln(out, ans)
	for i := range a {
		for _, x := range a[i] {
			fmt.Fprint(out, x, " ")
		}
		fmt.Fprintln(out)
	}
}
