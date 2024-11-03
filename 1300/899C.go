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
	if n&1 == 0 {
		if (n/2)&1 == 0 {
			fmt.Fprintln(out, 0)
			fmt.Fprint(out, n/2, " ")
			for i := 1; i < n/2; i += 2 {
				fmt.Fprint(out, i, n-i+1, " ")
			}
		} else {
			fmt.Fprintln(out, 1)
			fmt.Fprint(out, n/2, " ")
			for i := 1; i < n/2; i += 2 {
				fmt.Fprint(out, i, n-i+1, " ")
			}
			fmt.Fprint(out, n/2)
		}
	} else {
		n -= 3
		if (n/2)&1 == 0 {
			fmt.Fprintln(out, 0)
			fmt.Fprint(out, n/2+1, " ")
			fmt.Fprint(out, 3, " ")
			for i := 4; i < n/2+3; i += 2 {
				fmt.Fprint(out, i, n-i+7, " ")
			}
		} else {
			fmt.Fprintln(out, 1)
			fmt.Fprint(out, n/2+1, " ")
			fmt.Fprint(out, 3, " ")
			for i := 4; i < n/2+3; i += 2 {
				fmt.Fprint(out, i, n-i+7, " ")
			}
			fmt.Fprint(out, n/2+3)
		}
	}
}
