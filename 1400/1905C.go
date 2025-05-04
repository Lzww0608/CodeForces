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
	s := []byte{}
	fmt.Fscan(in, &s)

	f := true
	for i := 1; i < n; i++ {
		if s[i] < s[i-1] {
			f = false
			break
		}
	}
	if f {
		fmt.Fprintln(out, 0)
		return
	}

	st := []int{}
	for i := 0; i < n; i++ {
		for len(st) > 0 && s[st[len(st)-1]] < s[i] {
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}

	cnt := 0
	for _, i := range st {
		if s[st[0]] == s[i] {
			cnt++
		}
	}

	ans := len(st) - cnt
	for i, j := 0, len(st)-1; i < j; i, j = i+1, j-1 {
		s[st[i]], s[st[j]] = s[st[j]], s[st[i]]
	}

	f = true
	for i := 1; i < n; i++ {
		if s[i] < s[i-1] {
			f = false
			break
		}
	}
	if !f {
		fmt.Fprintln(out, -1)
		return
	}
	fmt.Fprintln(out, ans)
}
