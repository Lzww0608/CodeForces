package main

import (
	"bufio"
	"fmt"
	"os"
)

var cnt []int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for i, x := 3, 2; x <= int(1e12); i++ {
		x *= i
		cnt = append(cnt, x)
	}

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}

}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if (j-i)*2 == n {
				fmt.Fprint(out, 0, " ")
			} else if (j-i)*2 < n {
				fmt.Fprint(out, 1, " ")
			} else {
				fmt.Fprint(out, -1, " ")
			}
		}
	}
	fmt.Fprintln(out)
}
