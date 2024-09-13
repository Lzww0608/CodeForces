package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t, n int
	var s string
	for fmt.Scan(&t); t > 0; t-- {
		fmt.Fscan(in, &n, &s)
		cnt := 0
		for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
			if s[l] != s[r] {
				cnt++
			}
		}

		if n&1 == 0 {
			for i := 0; i <= n; i++ {
				if i >= cnt && i <= n-cnt && (i-cnt)%2 == 0 {
					fmt.Fprint(out, "1")
				} else {
					fmt.Fprint(out, "0")
				}
			}
			fmt.Fprintln(out)
		} else {
			fmt.Fprint(out, strings.Repeat("0", cnt))
			fmt.Fprint(out, strings.Repeat("1", n+1-cnt*2))
			fmt.Fprintln(out, strings.Repeat("0", cnt))
		}
	}
}
