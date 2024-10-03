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
		var s string
		fmt.Fscan(in, &s)
		ans := 0
		for i := 0; i < n; i++ {
			if s[i] == '0' {
				if i+1 < n && s[i+1] == '0' {
					ans += 2
				} else if i+2 < n && s[i+1] == '1' && s[i+2] == '0' {
					ans++
				}
			}

		}
		fmt.Fprintln(out, ans)
	}
}
