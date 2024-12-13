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
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n, p int
	fmt.Fscan(in, &n, &p)
	k := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			fmt.Fprintln(out, i+1, j+1)
			if k++; k == 2*n+p {
				return
			}
		}
	}
}
