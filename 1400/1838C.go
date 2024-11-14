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
	var n, m int
	fmt.Fscan(in, &n, &m)
	if check(m) {
		for i := 0; i < n; i++ {
			for j := 1; j <= m; j++ {
				fmt.Fprint(out, i*m+j, " ")
			}
			fmt.Fprintln(out)
		}
	} else {
		k, p := 0, n/2
		for i := 0; i < n; i++ {
			if i&1 == 1 {
				for j := 1; j <= m; j++ {
					fmt.Fprint(out, k*m+j, " ")
				}
				k++
			} else {
				for j := 1; j <= m; j++ {
					fmt.Fprint(out, p*m+j, " ")
				}
				p++
			}

			fmt.Fprintln(out)
		}

	}
	fmt.Fprintln(out)
}

func check(x int) bool {
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return true
		}
	}
	return false
}
