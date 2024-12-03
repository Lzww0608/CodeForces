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
	for i := 1; i < k-(n-k); i++ {
		fmt.Fprint(out, i, " ")
	}
	for i := k; i >= k-(n-k); i-- {
		fmt.Fprint(out, i, " ")
	}
	fmt.Fprintln(out)
}
