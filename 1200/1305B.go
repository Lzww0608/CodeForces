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
	l, r := 0, n-1
	cnt := 0
	a := []int{}
	b := []int{}
	for l < r {
		for l < r && s[l] != '(' {
			l++
		}
		for l < r && s[r] != ')' {
			r--
		}
		if l < r {
			cnt++
			a = append(a, l+1)
			b = append(b, r+1)
			l++
			r--
		}

	}
	if cnt == 0 {
		fmt.Fprintln(out, 0)
		return
	}
	fmt.Fprintln(out, 1)
	fmt.Fprintln(out, cnt<<1)
	for _, x := range a {
		fmt.Fprint(out, x, " ")
	}
	m := len(b)
	for i := m - 1; i >= 0; i-- {
		fmt.Fprint(out, b[i], " ")
	}
	fmt.Fprintln(out)
}
