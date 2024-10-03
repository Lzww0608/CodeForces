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
	sum, cnt := a[0], 0
	fmt.Fprint(out, a[0], " ")
	if a[0]&1 == 1 {
		cnt++
	}

	for i := 1; i < n; i++ {
		sum += a[i]
		if a[i]&1 == 1 {
			cnt++
		}
		cur := sum - cnt/3
		if cnt%3 == 1 {
			cur--
		}
		fmt.Fprint(out, cur, " ")
	}
	fmt.Fprintln(out)
}
