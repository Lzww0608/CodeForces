package main

import (
	"bufio"
	"bytes"
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
	var n, k int
	fmt.Fscan(in, &n, &k)
	k--
	s := make([]byte, n)
	fmt.Fscan(in, &s)

	t := make([]byte, n)
	for i := 0; i < k*2; i += 2 {
		t[i], t[i+1] = '(', ')'
	}
	for l, r := k*2, n-1; l < r; l, r = l+1, r-1 {
		t[l], t[r] = '(', ')'
	}

	ans := [][2]int{}
	for i := 0; i < n; i++ {
		if s[i] == t[i] {
			continue
		}
		j := i + 1
		for j < n && s[i] == s[j] {
			j++
		}
		ans = append(ans, [2]int{i, j})
		for l, r := i, j; l < r; l, r = l+1, r-1 {
			s[l], s[r] = s[r], s[l]
		}
	}

	if !bytes.Equal(s, t) {
		panic("not equal")
	}
	//fmt.Fprintln(out, string(s))
	fmt.Fprintln(out, len(ans))
	for _, v := range ans {
		fmt.Fprintln(out, v[0]+1, v[1]+1)
	}
}
