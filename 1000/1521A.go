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

	var t, a, b int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &a, &b)
		if b == 1 {
			fmt.Fprintln(out, "NO")
			continue
		}
		fmt.Fprintln(out, "YES")
		fmt.Fprintln(out, a*(b+1), a*(b-1), a*b*2)
	}
}
