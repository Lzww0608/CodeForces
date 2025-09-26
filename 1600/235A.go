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

	if n <= 2 {
		fmt.Fprintln(out, n)
	} else if n&1 == 1 {
		fmt.Fprintln(out, n*(n-1)*(n-2))
	} else if n%3 != 0 {
		fmt.Fprintln(out, n*(n-1)*(n-3))
	} else {
		fmt.Fprintln(out, (n-2)*(n-1)*(n-3))
	}
}
