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

	var n int
	fmt.Fscan(in, &n)
	var s string
	fmt.Fscan(in, &s)
	ans := 0
	for {
		pos, n := -1, len(s)
		for i := 0; i < n; i++ {
			if i > 0 {
				if (pos == -1 || s[i] > s[pos]) && s[i-1] == s[i]-1 {
					pos = i
				}
			}

			if i < n-1 {
				if (pos == -1 || s[i] > s[pos]) && s[i+1] == s[i]-1 {
					pos = i
				}
			}
		}
		if pos == -1 {
			break
		}
		s = s[:pos] + s[pos+1:]
		ans++
	}
	fmt.Fprintln(out, ans)
}
