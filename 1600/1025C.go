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
	s = s + s
	ans, cur := 1, 1
	pre := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] != pre {
			cur++
			ans = max(ans, cur)
		} else {
			cur = 1
		}
		pre = s[i]
	}
	fmt.Fprintln(out, min(ans, len(s)/2))
}
