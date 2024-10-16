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

	var s string
	var m int
	fmt.Fscan(in, &s)
	fmt.Fscan(in, &m)
	n := len(s)
	cnt := 0
	for i := range s {
		if s[i] == '?' || s[i] == '*' {
			cnt++
		}
	}
	if m < n-cnt*2 || (m > n-cnt && !strings.Contains(s, "*")) {
		fmt.Fprintln(out, "Impossible")
		return
	}
	ans := make([]byte, 0, m)
	k := n - m - cnt
	for i := range s {
		if s[i] == '?' {
			if k > 0 {
				ans = ans[:len(ans)-1]
				k--
			}
		} else if s[i] == '*' {
			if k > 0 {
				ans = ans[:len(ans)-1]
				k--
			} else if k < 0 {
				for k < 0 {
					k++
					ans = append(ans, ans[len(ans)-1])
				}
			}
		} else {
			ans = append(ans, s[i])
		}
	}
	fmt.Fprintln(out, string(ans))
}
