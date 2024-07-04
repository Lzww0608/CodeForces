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
		n--
		if n&1 == 1 {
			fmt.Fprintln(out, n/2, n/2+1, 1)
		} else if (n/2)&1 == 0 {
			fmt.Fprintln(out, n/2-1, n/2+1, 1)
		} else {
			fmt.Fprintln(out, n/2-2, n/2+2, 1)
		}

	}
}
