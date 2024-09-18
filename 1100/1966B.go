package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var q, n, m int
	for fmt.Fscan(in, &q); q > 0; q-- {
		fmt.Fscan(in, &n, &m)
		a := make([]string, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}
		if n == 1 && m == 1 {
			fmt.Fprintln(out, "YES")
			continue
		}
		if a[0] == strings.Repeat("W", m) && a[n-1] == strings.Repeat("B", m) {
			fmt.Fprintln(out, "NO")
			continue
		} else if a[n-1] == strings.Repeat("W", m) && a[0] == strings.Repeat("B", m) {
			fmt.Fprintln(out, "NO")
			continue
		} else {
			s := make([]byte, n)
			t := make([]byte, n)
			for i := range a {
				s[i] = a[i][0]
				t[i] = a[i][m-1]
			}
			if string(t) == strings.Repeat("W", n) && string(s) == strings.Repeat("B", n) {
				fmt.Fprintln(out, "NO")
				continue
			} else if string(s) == strings.Repeat("W", n) && string(t) == strings.Repeat("B", n) {
				fmt.Fprintln(out, "NO")
				continue
			}
		}
		fmt.Fprintln(out, "YES")
	}
}
