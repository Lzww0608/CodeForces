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

	var n int
	fmt.Fscan(in, &n)

	fmt.Fprintln(out, 7+(n-1)*3)
	fmt.Fprintln(out, 0, 0)
	fmt.Fprintln(out, 1, 0)
	for i := 1; i <= n; i++ {
		fmt.Fprintln(out, i-1, i)
		fmt.Fprintln(out, i, i)
		fmt.Fprintln(out, i+1, i)
	}
	fmt.Fprintln(out, n, n+1)
	fmt.Fprintln(out, n+1, n+1)
}
