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
	a := make([]int, n)
	b := make([][]int, 1001)
	for i := range a {
		fmt.Fscan(in, &a[i])
		b[calc(a[i])] = append(b[calc(a[i])], i)
	}
	cnt := 0
	ans := make([]int, n)
	for i := range b {
		if len(b[i]) > 0 {
			cnt++
			for _, x := range b[i] {
				ans[x] = cnt
			}
		}
	}

	fmt.Fprintln(out, cnt)
	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)
}

func calc(x int) int {
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return i
		}
	}
	return x
}
