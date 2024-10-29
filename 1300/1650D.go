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

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n, n*2)
	pos := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		a[i]--
		pos[a[i]] = i
	}
	ans := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for a[len(a)-1] != i {
			t := a[0]
			a = a[1:]
			a = append(a, t)
			ans[i]++
		}
		a = a[:len(a)-1]
	}
	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)
}
