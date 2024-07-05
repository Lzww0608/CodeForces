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
		fmt.Fprintln(out, (n+1)>>1)
		for i := 1; i < 3*n+1-i; i += 3 {
			fmt.Fprintln(out, i, " ", 3*n+1-i)
		}

	}
}
