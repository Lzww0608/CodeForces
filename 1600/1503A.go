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

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	var a string
	fmt.Fscan(in, &a)
	cnt := strings.Count(a, "1")
	if cnt&1 == 1 || a[0] != '1' || a[n-1] != '1' {
		fmt.Fprintln(out, "NO")
		return
	}

	s := make([]byte, n)
	t := make([]byte, n)
	cur1, cur2 := 0, 0
	for i := 0; i < n; i++ {
		if a[i] == '1' {
			if cur1 < cnt/2 {
				s[i] = '('
				t[i] = '('
			} else {
				s[i] = ')'
				t[i] = ')'
			}
			cur1++
		} else {
			if cur2&1 == 0 {
				s[i] = '('
				t[i] = ')'
			} else {
				s[i] = ')'
				t[i] = '('
			}
			cur2++
		}

	}
	//fmt.Fprintln(out, string(t))

	cur := 0
	for i := 0; i < n; i++ {
		if t[i] == '(' {
			cur++
		} else {
			cur--
		}
		if cur < 0 {
			fmt.Fprintln(out, "NO")
			return
		}
	}

	if cur != 0 {
		fmt.Fprintln(out, "NO")
		return
	}

	fmt.Fprintln(out, "YES")
	fmt.Fprintln(out, string(s))
	fmt.Fprintln(out, string(t))
}
