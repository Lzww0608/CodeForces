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
		if (x+y)&1 == 1 {
			fmt.Fprintln(out, -1, " ", -1)
			continue
		}
		if x&1 == 1 {
			if x > y {
				fmt.Fprintln(out, x/2, " ", y/2+1)
			} else {
				fmt.Fprintln(out, x/2+1, " ", y/2)
			}

		} else {
			fmt.Fprintln(out, x/2, " ", y/2)
		}
	}
}
