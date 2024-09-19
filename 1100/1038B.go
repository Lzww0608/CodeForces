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
	if n < 3 {
		fmt.Fprintln(out, "No")
	} else {
		fmt.Fprintln(out, "Yes")
		fmt.Fprint(out, n-n/2, " ")
		for i := 1; i <= n; i += 2 {
			fmt.Fprint(out, i, " ")
		}
		fmt.Fprintln(out)
		fmt.Fprint(out, n/2, " ")
		for i := 2; i <= n; i += 2 {
			fmt.Fprint(out, i, " ")
		}
		fmt.Fprintln(out)
	}
}
