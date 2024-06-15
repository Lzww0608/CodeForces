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

	var t, n, m int
next:
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n, &m)
		var s, t string
		fmt.Fscan(in, &s)
		fmt.Fscan(in, &t)
		zero := 0
		for i := 0; i < n-1; i++ {
			if s[i] == s[i+1] {
				if s[i] == '0' {
					if zero < 0 {
						fmt.Fprintln(out, "NO")
						continue next
					}
					zero++
				} else if zero > 0 {
					fmt.Fprintln(out, "NO")
					continue next
				} else {
					zero--
				}
			}
		}

		if zero == 0 {
			fmt.Fprintln(out, "YES")
			continue next
		}
		if (zero > 0 && (t[0] != '1' || t[m-1] != '1')) || (zero < 0 && (t[0] != '0' || t[m-1] != '0')) {
			fmt.Fprintln(out, "NO")
			continue next
		}

		for i := 0; i < m-1; i++ {
			if t[i] == t[i+1] {
				fmt.Fprintln(out, "NO")
				continue next
			}
		}
		fmt.Fprintln(out, "YES")
	}
}
