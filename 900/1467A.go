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
		fmt.Fprint(out, 9)
		for i := 1; i < n; i++ {
			fmt.Fprint(out, (7+i)%10)
		}
		fmt.Fprintln(out)
	}
}
