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
	var s []byte
	fmt.Fscan(in, &s)
	n := len(s)

	for i := range n {
		l, r := i, min(n-1, i+9)

		p, mx := i, 0
		for j := l; j <= r; j++ {
			x := int(s[j] - '0')
			if x-(j-l) > mx {
				mx = x - (j - l)
				p = j
			}
		}

		s[p] = byte('0' + (int(s[p]-'0') - (p - i)))
		for j := p; j > i; j-- {
			s[j], s[j-1] = s[j-1], s[j]
		}
	}

	fmt.Fprintln(out, string(s))
}
