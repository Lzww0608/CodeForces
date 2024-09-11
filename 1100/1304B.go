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

	var n, m int
	ans := 0
	var left, right string
	fmt.Fscan(in, &n, &m)
	s := make([]string, n)
	cnt := map[string]int{}
	for i := range s {
		fmt.Fscan(in, &s[i])
		tmp := []byte(s[i])
		for l, r := 0, m-1; l < r; l, r = l+1, r-1 {
			tmp[l], tmp[r] = tmp[r], tmp[l]
		}
		if cnt[string(tmp)] > 0 {
			cnt[string(tmp)]--
			ans += m * 2
			left += s[i]
			right = string(tmp) + right
		} else {
			cnt[s[i]]++
		}
	}

	for _, t := range s {
		if cnt[t] > 0 {
			f := true
			for l, r := 0, m-1; l < r; l, r = l+1, r-1 {
				if t[l] != t[r] {
					f = false
					break
				}
			}
			if f {
				ans += m
				left += t
				break
			}
		}
	}

	fmt.Fprintln(out, ans)
	fmt.Fprintln(out, left+right)
}
