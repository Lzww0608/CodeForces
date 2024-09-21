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

	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if (i+j)&1 == 0 {
				fmt.Fprint(out, "W")
			} else {
				fmt.Fprint(out, "B")
			}
		}
		fmt.Fprintln(out)
	}
}
