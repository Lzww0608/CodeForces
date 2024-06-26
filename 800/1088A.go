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

	var x int
	fmt.Fscan(in, &x)
	for b := x; b*b >= x; b-- {
		for a := b; a > 0; a-- {
			if b%a == 0 && a*b > x && a/b < x {
				fmt.Fprintln(out, b, a)
				return
			}
		}
	}

	fmt.Fprintln(out, -1)
}
