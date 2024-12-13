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

	var n, k int
	fmt.Fscan(in, &n, &k)
	var s string
	fmt.Fscan(in, &s)
	t := make([]byte, k)

	cnt := map[byte]bool{}
	mn := byte('z')
	for i := 0; i < n; i++ {
		cnt[s[i]] = true
		mn = min(mn, s[i])
	}

	for i := 0; i < k; i++ {
		if i < n {
			t[i] = s[i]
		} else {
			t[i] = mn
		}
	}

	if k > n {
		fmt.Fprintln(out, string(t))
		return
	}

	for i := k - 1; i >= 0; i-- {
		c := t[i] + 1
		for c < 'z' && !cnt[c] {
			c++
		}

		if c <= 'z' && cnt[c] {
			t[i] = c
			break
		}
		t[i] = mn
	}
	fmt.Fprintln(out, string(t))
}
