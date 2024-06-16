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
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j == i {
				fmt.Fprint(out, k, " ")
			} else {
				fmt.Fprint(out, 0, " ")
			}
		}
		fmt.Fprintln(out)
	}
}
