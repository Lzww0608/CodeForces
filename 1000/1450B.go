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
	var n, k int
	fmt.Fscan(in, &n, &k)
	a := make([][2]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i][0], &a[i][1])
	}

next:
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if abs(a[i][0]-a[j][0])+abs(a[i][1]-a[j][1]) > k {
				continue next
			}
		}
		fmt.Fprintln(out, 1)
		return
	}
	fmt.Fprintln(out, -1)
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
