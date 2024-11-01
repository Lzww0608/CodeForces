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
	var n, k int
	fmt.Fscan(in, &n, &k)
	t := k
	var str string
	fmt.Fscan(in, &str)
	s := []byte(str)
	f := make([]int, n)
	for i := 0; i < n && k > 0; i++ {
		if t&1 == int(s[i]-'0') {
			f[i] = 1
			k--
		}
	}
	f[n-1] += k
	for i := 0; i < n; i++ {
		if (t-f[i])&1 == 1 {
			if s[i] == '1' {
				s[i] = '0'
			} else {
				s[i] = '1'
			}
		}
	}

	fmt.Fprintln(out, string(s))
	for _, x := range f {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)
}
