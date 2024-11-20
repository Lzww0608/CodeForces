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

	var n, m int
	fmt.Fscan(in, &n, &m)
	if n == 1 && m == 1 {
		fmt.Fprintln(out, 0)
		return
	}

	if m == 1 {
		for i := 1; i <= n; i++ {
			fmt.Fprint(out, i+1, " ")
		}
		fmt.Fprintln(out)
	} else {
		for i := 1; i <= n; i++ {
			for j := n + 1; j <= n+m; j++ {
				fmt.Fprint(out, i*j, " ")
			}
			fmt.Fprintln(out)
		}
	}

}
