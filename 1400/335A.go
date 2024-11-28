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
	fmt.Fscan(in, &s)
	var n int
	fmt.Fscan(in, &n)

	m := map[byte]int{}
	for i := range s {
		m[s[i]]++
	}

	if len(m) > n {
		fmt.Fprintln(out, -1)
		return
	}

	check := func(t int) bool {
		cnt := 0
		for _, v := range m {
			cnt += (v + t - 1) / t
		}
		if cnt > n {
			return false
		}
		return true
	}

	l, r := 1, 1001
	ans := 1
	for l < r {
		mid := l + ((r - l) >> 1)
		if check(mid) {
			ans = mid
			r = mid
		} else {
			l = mid + 1
		}
	}

	fmt.Fprintln(out, ans)
	sum := 0
	for c, x := range m {
		cnt := (x + ans - 1) / ans
		sum += cnt
		fmt.Fprint(out, strings.Repeat(string(c), cnt))
	}
	if sum < n {
		fmt.Fprintln(out, strings.Repeat("a", n-sum))
	}

}
