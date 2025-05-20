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
	sum := 0

	mx := 0
	for i := range a {
		fmt.Fscan(in, &a[i])
		sum += a[i]
		mx = max(mx, a[i])
	}
	if mx > sum/2 || sum&1 == 1 {
		fmt.Fprintln(out, "T")
		return
	}

	fmt.Fprintln(out, "HL")
}
