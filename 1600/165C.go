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

	if k == 0 {
		ans := 0
		cnt := 0
		for i := range s {
			if s[i] == '1' {
				cnt = 0
			} else {
				cnt++
				ans += cnt
			}
		}

		fmt.Fprintln(out, ans)
		return
	}

	ans := 0
	cnt := 0
	l, r := 0, 0
	for r < len(s) && cnt < k {
		if s[r] == '1' {
			cnt++
		}
		r++
	}
	if cnt < k {
		fmt.Fprintln(out, 0)
		return
	}

	for {
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

		if r == len(s) {
			break
		}
		l++
		r++
	}

	fmt.Fprintln(out, ans)
}
