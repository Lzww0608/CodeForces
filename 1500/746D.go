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

	var n, k, a, b int
	fmt.Fscan(in, &n, &k, &a, &b)
	if n/(k+1) > b || n/(k+1) > a {
		fmt.Fprintln(out, "NO")
		return
	}

	pre := byte(' ')
	for a > 0 || b > 0 {
		if a > b {
			if pre == 'G' {
				b--
				fmt.Fprint(out, "B")
			}
			t := min(a-b, k)
			fmt.Fprint(out, strings.Repeat("G", t))
			a -= t
			pre = 'G'
		} else if b > a {
			if pre == 'B' {
				a--
				fmt.Fprint(out, "G")
			}
			t := min(b-a, k)
			fmt.Fprint(out, strings.Repeat("B", t))
			b -= t
			pre = 'B'
		} else {
			if pre == 'B' {
				fmt.Fprint(out, "GB")
			} else {
				fmt.Fprint(out, "BG")
			}
			a--
			b--
		}
	}
}
