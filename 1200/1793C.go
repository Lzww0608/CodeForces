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

	var t, n int
next:
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}

		l, r := 0, n-1
		mx, mn := n, 1
		for r-l >= 3 {
			if a[l] != mx && a[l] != mn && a[r] != mx && a[r] != mn {
				fmt.Fprintln(out, l+1, r+1)
				continue next
			}
			if a[l] == mn {
				if a[r] == mx {
					mn++
					mx--
					l++
					r--
				} else {
					mn++
					l++
				}
			} else if a[l] == mx {
				if a[r] == mn {
					mn++
					mx--
					l++
					r--
				} else {
					mx--
					l++
				}
			} else if a[r] == mx {
				mx--
				r--
			} else if a[r] == mn {
				mn++
				r--
			}

		}
		fmt.Fprintln(out, -1)
	}
}
