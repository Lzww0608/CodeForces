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
	var n, x int
	fmt.Fscan(in, &n, &x)
	a := make([]int, n)
	cnt := 0
	sum := 0
	for i := range a {
		fmt.Fscan(in, &a[i])
		sum += a[i]
		if a[i] == x {
			cnt++
		}
	}

	if cnt == n {
		fmt.Fprintln(out, 0)
	} else if cnt > 0 || sum == n*x {
		fmt.Fprintln(out, 1)
	} else {
		fmt.Fprintln(out, 2)
	}

}
