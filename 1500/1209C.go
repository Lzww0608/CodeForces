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
	var n int
	fmt.Fscan(in, &n)
	var s string
	fmt.Fscan(in, &s)

next:
	for x := 1; x < 9; x++ {
		a := make([]int, 0, n)
		b := make([]int, 0, n)
		for i := 0; i < n; i++ {
			y := int(s[i] - '0')
			if y < x {
				if len(a) == 0 || s[a[len(a)-1]] <= s[i] {
					a = append(a, i)
				} else {
					continue next
				}
			} else if y > x {
				if len(b) == 0 || s[b[len(b)-1]] <= s[i] {
					b = append(b, i)
				} else {
					continue next
				}
			}
		}
		for i := 0; i < n; i++ {
			y := int(s[i] - '0')
			if x == y {
				if len(a) == 0 || a[len(a)-1] < i || len(b) == 0 || b[0] > i {

				} else {
					continue next
				}
			}
		}

		for i := 0; i < n; i++ {
			y := int(s[i] - '0')
			if y < x {
				fmt.Fprint(out, 1)
			} else if y > x {
				fmt.Fprint(out, 2)
			} else {
				if len(a) == 0 || a[len(a)-1] < i {
					fmt.Fprint(out, 1)
				} else {
					fmt.Fprint(out, 2)
				}
			}
		}
		fmt.Fprintln(out)
		return
	}
	fmt.Fprintln(out, "-")
}
