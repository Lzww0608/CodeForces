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

	var t, n, k int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n, &k)
		for i := 0; i < k; i++ {
			fmt.Fprint(out, "a")
		}
		for i := k; i < n; i++ {
			fmt.Fprint(out, string('a'+(i-k+1)%3))
		}
		fmt.Fprintln(out)
	}
}
