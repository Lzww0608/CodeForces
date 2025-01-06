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

	fmt.Fprintln(out, n+n/2)
	if n&1 == 0 {
		for i := 1; i <= n; i += 2 {
			fmt.Fprint(out, i, " ")
		}
		for i := 2; i <= n; i += 2 {
			fmt.Fprint(out, i, " ")
		}
		for i := 1; i <= n; i += 2 {
			fmt.Fprint(out, i, " ")
		}
	} else {
		for i := 2; i <= n; i += 2 {
			fmt.Fprint(out, i, " ")
		}
		for i := 1; i <= n; i += 2 {
			fmt.Fprint(out, i, " ")
		}
		for i := 2; i <= n; i += 2 {
			fmt.Fprint(out, i, " ")
		}
	}

}
