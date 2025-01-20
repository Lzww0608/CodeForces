package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const N int = 26

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s, t string
	fmt.Fscan(in, &s, &t)
	m, n := len(s), len(t)
	if n > m {
		s = strings.Replace(s, "?", "a", -1)
		fmt.Fprintln(out, s)
		return
	}
	k := strings.Count(s, "?")
	cntS := make(map[byte]int)
	cntT := make(map[byte]int)
	for i := range s {
		cntS[s[i]]++
	}
	for i := range t {
		cntT[t[i]]++
	}

	check := func(mid int) bool {
		cost := 0
		for i := range cntT {
			cost += max(0, cntT[i]*mid-cntS[i])
		}
		return cost <= k
	}

	l, r := -1, m/n+1
	for l+1 < r {
		mid := l + ((r - l) >> 1)
		if check(mid) {
			l = mid
		} else {
			r = mid
		}
	}

	ans := make([]byte, 0, m)
	bytes := make(map[byte]int)
	for i := 0; i < N; i++ {
		b := byte('a' + i)
		bytes[b] = max(0, cntT[b]*l-cntS[b])
		k -= bytes[b]
	}

	var j byte = 'a'
	for i := 0; i < m; i++ {
		if s[i] != '?' {
			ans = append(ans, s[i])
		} else {
			if k > 0 {
				ans = append(ans, 'a')
				k--
			} else {
				for bytes[j] == 0 {
					j++
				}
				bytes[j]--
				ans = append(ans, j)
			}
		}
	}
	fmt.Fprintln(out, string(ans))
}
