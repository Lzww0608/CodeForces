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

	var n, k int
	fmt.Fscan(in, &n, &k)
	if n == 1 {
		if k != 0 {
			fmt.Fprintln(out, -1)
		} else {
			fmt.Fprintln(out, 1)
		}
		return
	}
	if k < n/2 {
		fmt.Fprintln(out, -1)
		return
	}
	x := k - (n-2)/2
	fmt.Fprint(out, x, 2*x, " ")
	for i := 1; i <= n-2; i++ {
		fmt.Fprint(out, 2*x+i, " ")
	}

}
