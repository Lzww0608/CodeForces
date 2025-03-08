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

	var n1, m1, n2, m2, n3, m3 int
	fmt.Fscan(in, &n1, &m1, &n2, &m2, &n3, &m3)
	n1, m1 = min(n1, m1), max(n1, m1)
	n2, m2 = min(n2, m2), max(n2, m2)
	n3, m3 = min(n3, m3), max(n3, m3)
	if m1 == m2 && m2 == m3 && n1+n2+n3 == m3 {
		n := m3
		fmt.Fprintln(out, n)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if j < n1 {
					fmt.Fprint(out, "A")
				} else if j < n1+n2 {
					fmt.Fprint(out, "B")
				} else {
					fmt.Fprint(out, "C")
				}
			}
			fmt.Fprintln(out)
		}
		return
	}

	n := max(m1, m2, m3)
	if n == m1 {
		solve(out, n1, m1, n2, m2, n3, m3, "A", "B", "C")
	} else if n == m2 {
		solve(out, n2, m2, n1, m1, n3, m3, "B", "A", "C")
	} else {
		solve(out, n3, m3, n1, m1, n2, m2, "C", "A", "B")
	}
}

func solve(out *bufio.Writer, n1, m1, n2, m2, n3, m3 int, c1, c2, c3 string) {
	n := m1
	d := n - n1
	if n2 != d && m2 != d || n3 != d && m3 != d || n2+m2+n3+m3 != d*2+n {
		fmt.Fprintln(out, -1)
		return
	}
	fmt.Fprintln(out, n)
	for i := 0; i < n1; i++ {
		for j := 0; j < n; j++ {
			fmt.Fprint(out, c1)
		}
		fmt.Fprintln(out)
	}

	m := n2 + m2 - d
	for i := n1; i < n; i++ {
		for j := 0; j < n; j++ {
			if j < m {
				fmt.Fprint(out, c2)
			} else {
				fmt.Fprint(out, c3)
			}
		}
		fmt.Fprintln(out)
	}
}
