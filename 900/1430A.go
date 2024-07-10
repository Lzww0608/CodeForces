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

	var t, n int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		if n < 5 && n != 3 {
			fmt.Fprintln(out, -1)
			continue
		}
		x := n % 3
		if x == 0 {
			fmt.Fprintln(out, n/3, 0, 0)
		} else if x == 1 {
			fmt.Fprintln(out, n/3-2, 0, 1)
		} else {
			fmt.Fprintln(out, n/3-1, 1, 0)
		}
	}
}
