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
	if n&1 == 1 {
		fmt.Fprintln(out, -1)
		return
	}
	ans := [][]int{}
	for i := 0; i < n; i += 2 {
		if a[i] == a[i+1] {
			ans = append(ans, []int{i + 1, i + 2})
		} else {
			ans = append(ans, []int{i + 1, i + 1})
			ans = append(ans, []int{i + 2, i + 2})
		}
	}
	fmt.Fprintln(out, len(ans))
	for _, v := range ans {
		fmt.Fprintln(out, v[0], v[1])
	}
}
