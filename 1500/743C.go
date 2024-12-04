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
	if n == 1 {
		fmt.Fprintln(out, -1)
		return
	}
	fmt.Fprintln(out, n, n+1, n*(n+1))
}
