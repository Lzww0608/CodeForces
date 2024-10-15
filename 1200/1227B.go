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
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	mx := 0
	b := make([]int, n)
	vis := make([]bool, n+1)
	for i, x := range a {
		if x < mx || x <= i {
			fmt.Fprintln(out, -1)
			return
		}
		if x > mx {
			mx = x
			b[i] = x
		}
		vis[x] = true
	}
	j := 1
	for i := range b {
		for j <= n && vis[j] {
			j++
		}
		if b[i] == 0 {
			b[i] = j
			j++
		}
	}
	for _, x := range b {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)
}
