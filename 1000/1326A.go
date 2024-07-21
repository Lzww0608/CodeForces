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

	var t, n int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		if n == 1 {
			fmt.Fprintln(out, -1)
			continue
		}
		fmt.Fprint(out, 2)
		for i := 0; i < n-1; i++ {
			fmt.Fprint(out, 3)
		}
		fmt.Fprintln(out)
	}
}
