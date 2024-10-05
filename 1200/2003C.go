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

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	var s string
	fmt.Fscan(in, &s)
	cnt := make([]int, 26)
	for i := range s {
		x := int(s[i]&31) - 1
		cnt[x]++
	}
	ans := []byte{}
	for i := 0; len(ans) < n; i++ {
		if cnt[i%26] > 0 {
			cnt[i%26]--
			ans = append(ans, byte('a'+i%26))
		}
	}
	fmt.Fprintln(out, string(ans))
}
