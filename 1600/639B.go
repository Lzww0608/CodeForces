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

	var n, d, h int
	fmt.Fscan(in, &n, &d, &h)
	if d > 2*h || d < h || (d == 1 && h == 1 && n > 2) {
		fmt.Fprintln(out, -1)
		return
	}

	i := 1
	for i = 1; i <= h; i++ {
		fmt.Fprintln(out, i, i+1)
	}
	if d > h {
		fmt.Fprintln(out, 1, i+1)
		i++
		for ; i <= d; i++ {
			fmt.Fprintln(out, i, i+1)
		}
		for ; i < n; i++ {
			fmt.Fprintln(out, 1, i+1)
		}
	} else {
		for ; i < n; i++ {
			fmt.Fprintln(out, 2, i+1)
		}
	}

}
