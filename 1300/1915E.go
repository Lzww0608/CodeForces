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
		if i&1 == 1 {
			a[i] = -a[i]
		}
	}
	vis := make(map[int]bool)
	vis[0] = true
	sum := 0
	for _, x := range a {
		sum += x
		if vis[sum] {
			fmt.Fprintln(out, "YES")
			return
		}
		vis[sum] = true
	}

	fmt.Fprintln(out, "NO")
}
