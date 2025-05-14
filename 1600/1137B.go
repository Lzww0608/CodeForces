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

	var s, t string
	fmt.Fscan(in, &s, &t)
	if len(t) > len(s) {
		fmt.Fprintln(out, s)
		return
	}

	cnt := [2]int{}
	for i := range s {
		cnt[int(s[i]-'0')]++
	}

	pi := make([]int, len(t))
	j := 0
	for i := 1; i < len(t); i++ {
		for j > 0 && t[j] != t[i] {
			j = pi[j-1]
		}

		if t[j] == t[i] {
			j++
		}
		pi[i] = j
	}

	ans := make([]byte, 0, len(s))
	j = 0
	for cnt[int(t[j]-'0')] > 0 {
		cnt[int(t[j]-'0')]--
		ans = append(ans, t[j])
		j++
		if j == len(t) {
			j = pi[j-1]
		}
	}

	fmt.Fprintf(out, "%s%s%s", ans, strings.Repeat("0", cnt[0]), strings.Repeat("1", cnt[1]))
}
