package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {

	var q int
	in := bufio.NewReader(os.Stdin)
	Fscan(in, &q)

	var a, b int
	var s string
	for ; q > 0; q-- {

		Fscan(in, &a, &b)
		Fscan(in, &s)

		ans := 0
		for i, last := 0, -1; i < len(s); i++ {
			if s[i] == '1' {
				if last != -1 {
					ans += min(a, (i-last-1)*b)
				} else {
					ans += a
				}
				last = i
			}
		}

		Println(ans)
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
