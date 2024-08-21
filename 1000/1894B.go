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
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		a := make([]int, n)
		b := make([]int, n)
		cnt := 0
		m := map[int]int{}
		for i := range a {
			fmt.Fscan(in, &a[i])
			b[i] = 1
			if m[a[i]] == 1 {
				cnt++
			}
			m[a[i]]++
		}
		if cnt < 2 {
			fmt.Fprintln(out, -1)
			continue
		}
		cnt = 0
		for i := range a {
			if m[a[i]] == 2 {
				b[i] = 2 + cnt
				cnt ^= 1
			}
			m[a[i]]--
			fmt.Fprint(out, b[i], " ")
		}
		fmt.Fprintln(out)
	}
}
