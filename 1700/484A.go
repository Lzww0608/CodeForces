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

	var t, l, r int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &l, &r)
		for l|(l+1) <= r {
			l |= l + 1
		}
		fmt.Fprintln(out, l)
	}
}
