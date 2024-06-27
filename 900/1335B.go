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

	var t, n, a, b int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n, &a, &b)
		for i := 0; i < n; i++ {
			fmt.Fprint(out, string('a'+i%b))
		}
		fmt.Fprintln(out)
	}
}
