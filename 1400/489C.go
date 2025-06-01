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

	var m, s int
	fmt.Fscan(in, &m, &s)
	if s > m*9 || m > 1 && s == 0 {
		fmt.Fprintln(out, -1, -1)
		return
	} else if m == 1 {
		fmt.Fprintln(out, s, s)
		return
	}

	a := make([]byte, 0, m)
	b := make([]byte, 0, m)
	for i, cur := 0, s; i < m; i++ {
		t := min(cur, 9)
		a = append(a, byte(t+'0'))
		cur -= t
	}

	for i, cur := 1, s; i <= m; i++ {
		if i == 1 {
			t := max(1, cur-(m-i)*9)
			cur -= t
			b = append(b, byte(t+'0'))
		} else {
			t := max(0, cur-(m-i)*9)
			cur -= t
			b = append(b, byte(t+'0'))
		}
	}

	fmt.Fprintln(out, string(b), string(a))
}
