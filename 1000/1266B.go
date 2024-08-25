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

	var t, x int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &x)
		if x <= 14 || x%14 > 6 || x%14 == 0 {
			fmt.Fprintln(out, "NO")
		} else {
			fmt.Fprintln(out, "YES")
		}

	}
}
