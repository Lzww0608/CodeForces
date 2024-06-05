package main

import (
	"bufio"
	"fmt"
	"os"
)

var MOD int = 1e9 + 7

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s string
	fmt.Fscan(in, &s)
	n := len(s)
	cnt, ans := 0, 0
	for i := n - 1; i >= 0; i-- {
		if s[i] == 'b' {
			cnt++
		} else {
			ans += cnt
			cnt <<= 1
			ans %= MOD
		}
		cnt %= MOD
	}
	fmt.Fprintln(out, ans)
}
