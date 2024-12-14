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

	if n < 6 {
		fmt.Fprintln(out, -1)
	} else {
		fmt.Fprintln(out, 1, 2)
		fmt.Fprintln(out, 1, 3)
		fmt.Fprintln(out, 1, 4)
		fmt.Fprintln(out, 2, 5)
		fmt.Fprintln(out, 2, 6)
		for i := 7; i <= n; i++ {
			fmt.Fprintln(out, 1, i)
		}
	}

	for i := 2; i <= n; i++ {
		fmt.Fprintln(out, 1, i)
	}
}
