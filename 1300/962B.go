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

	var n, a, b int
	fmt.Fscan(in, &n, &a, &b)
	var s string
	fmt.Fscan(in, &s)
	cnt, ans := 0, 0
	for i := 0; i <= n && (a > 0 || b > 0); i++ {
		if i < n && s[i] == '.' {
			cnt++
		} else if cnt > 0 {
			if cnt&1 == 0 {
				cnt >>= 1
				ans += min(a, cnt) + min(b, cnt)
				a -= min(a, cnt)
				b -= min(b, cnt)
			} else {
				x, y := cnt/2, cnt/2+1
				if a >= b {
					x, y = cnt/2+1, cnt/2
				}
				ans += min(a, x) + min(b, y)
				a -= min(a, x)
				b -= min(b, y)
			}
			cnt = 0
		}
	}
	fmt.Fprintln(out, ans)
}
