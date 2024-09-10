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
		ans := make([]byte, n)
		if s[0] == '9' {
			subtract := uint8(0)
			for i := n - 1; i >= 0; i-- {
				if s[i]+subtract <= '1' {
					ans[i] = byte('1' - s[i] - subtract + '0')
					subtract = 0
				} else {
					ans[i] = byte('1' + 10 - subtract - s[i] + '0')
					subtract = 1
				}
			}

		} else {
			for i := range ans {
				ans[i] = byte('0' + int('9'-s[i]))
			}
		}
		fmt.Fprintln(out, string(ans))
	}
}
