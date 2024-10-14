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

	var s string
	fmt.Fscan(in, &s)
	n := len(s)
	a := make([]int, 0, n)
	b := make([]int, 0, n)
	for i := n - 1; i >= 0; i-- {
		if s[i] == 'l' {
			a = append(a, i+1)
		} else {
			b = append(b, i+1)
		}
	}
	m := len(b)
	for i := m - 1; i >= 0; i-- {
		fmt.Fprintln(out, b[i])
	}

	for _, x := range a {
		fmt.Fprintln(out, x)
	}
	return
}
