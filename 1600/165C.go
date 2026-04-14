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

	var k int
	fmt.Fscan(in, &k)
	var s string
	fmt.Fscan(in, &s)

	ans := 0
	l, r := 0, 0
	for cnt := 0; r < len(s) && cnt < k; r++ {
		if s[r] == '1' {
			cnt++
		}
	}
	if r == len(s) {
		fmt.Fprintln(out, 0)
		return
	}

	for r < len(s) {
		x, y := 0, 0
		for l < len(s) && s[l] != '1' {
			l++
			x++
		}
		for r < len(s) && s[r] != '1' {
			r++
			y++
		}

		ans += (x + 1) * (y + 1)
		l++
		r++
	}

	fmt.Fprintln(out, ans)
}
