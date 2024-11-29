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

	ans, cur := 0, 0
	for i := 0; i < n; i++ {
		if a[i] != i+1 {
			if cur == 0 {
				ans++
			}
			cur++
		} else {
			cur = 0
		}

	}

	fmt.Fprintln(out, min(2, ans))
}
