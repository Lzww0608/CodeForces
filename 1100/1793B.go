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

	var t, x, y int

	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &x, &y)
		fmt.Fprintln(out, (x-y)*2)
		for i := x; i > y; i-- {
			fmt.Fprint(out, i, " ")
		}
		for i := y; i < x; i++ {
			fmt.Fprint(out, i, " ")
		}
		fmt.Fprintln(out)
	}
}
