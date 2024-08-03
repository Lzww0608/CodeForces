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

	var t int
	var a, b string
next:
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &a)
		fmt.Fscan(in, &b)
		n := len(a)

		for i := 0; i < n-1; i++ {
			if a[i] == b[i] && a[i] == '0' && a[i+1] == b[i+1] && a[i+1] == '1' {
				fmt.Fprintln(out, "YES")
				continue next
			}
		}
		fmt.Fprintln(out, "NO")
	}
}
