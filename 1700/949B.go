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

	var n, q, x int
	fmt.Fscan(in, &n, &q)
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &x)
		for x&1 != 1 {
			x = n + x/2
		}
		fmt.Fprintln(out, (x+1)>>1)
	}
}
