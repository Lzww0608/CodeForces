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

	var (
		t int
		s string
	)

	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &s)
		n := len(s)
		if s == "()" {
			fmt.Fprintln(out, "NO")
			continue
		}
		fmt.Fprintln(out, "YES")
		f := true
		for i := 1; i < n; i++ {
			if s[i] == s[i-1] {
				f = false
				break
			}
		}

		if !f {
			for i := 0; i < n; i++ {
				fmt.Fprint(out, "()")
			}
		} else {
			for i := 0; i < n; i++ {
				fmt.Fprint(out, "(")
			}
			for i := 0; i < n; i++ {
				fmt.Fprint(out, ")")
			}
		}
		fmt.Fprintln(out)
	}

}
