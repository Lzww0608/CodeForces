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

	var n, k int
	fmt.Fscan(in, &n, &k)
	var s string
	fmt.Fscan(in, &s)

	ans := 0
	a, b := 0, 0
	for l, r := 0, 0; r < n; r++ {
		if s[r] == 'a' {
			a++
		} else {
			b++
		}

		for min(a, b) > k {
			if s[l] == 'a' {
				a--
			} else {
				b--
			}
			l++
		}

		ans = max(ans, r-l+1)
	}

	fmt.Fprintln(out, ans)
}
