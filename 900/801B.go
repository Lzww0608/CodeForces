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

	var x, y string

	fmt.Fscan(in, &x)
	fmt.Fscan(in, &y)

	for i := range x {
		if x[i] < y[i] {
			fmt.Fprintln(out, -1)
			return
		}
	}
	fmt.Fprintln(out, y)
}
