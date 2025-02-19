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
	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	if k < n+m-2 || k&1 != (n+m-2)&1 {
		fmt.Fprintln(out, "NO")
		return
	}

	fmt.Fprintln(out, "YES")
	for i := 0; i < n; i++ {
		for j := 0; j < m-1; j++ {
			if i == 0 {
				if j&1 == 0 {
					fmt.Fprint(out, "R ")
				} else {
					fmt.Fprint(out, "B ")
				}
			} else if i == 1 && j == 0 {
				fmt.Fprint(out, "R ")
			} else if (i == n-1 || i == n-2) && j == m-2 {
				if (m-1)&1 == 1 {
					fmt.Fprint(out, "R ")
				} else {
					fmt.Fprint(out, "B ")
				}
			} else {
				fmt.Fprint(out, "R ")
			}
		}
		fmt.Fprintln(out)
	}
	for i := 0; i < n-1; i++ {
		for j := 0; j < m; j++ {
			if i == 0 {
				if j == 0 || j == 1 {
					fmt.Fprint(out, "B ")
				} else if j == m-1 {
					if m&1 == 0 {
						fmt.Fprint(out, "B ")
					} else {
						fmt.Fprint(out, "R ")
					}
				} else {
					fmt.Fprint(out, "B ")
				}
			} else if j == m-1 {
				if (i+m&1)&1 == 0 {
					fmt.Fprint(out, "B ")
				} else {
					fmt.Fprint(out, "R ")
				}
			} else {
				fmt.Fprint(out, "B ")
			}
		}
		fmt.Fprintln(out)
	}

}
