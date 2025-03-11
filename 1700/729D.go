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

	var n, a, b, k int
	fmt.Fscan(in, &n, &a, &b, &k)
	var s string
	fmt.Fscan(in, &s)
	ans, cnt := 0, 0
	p := []int{}
	for i := 0; i < n; i++ {
		if s[i] == '0' {
			if cnt++; cnt == b {
				ans++
				cnt = 0
				p = append(p, i+1)
			}
		} else {
			cnt = 0
		}
	}

	ans -= a - 1
	fmt.Fprintln(out, ans)
	for i := 0; i < ans; i++ {
		fmt.Fprint(out, p[i], " ")
	}
}
