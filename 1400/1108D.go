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

	var n int
	fmt.Fscan(in, &n)
	var t string
	fmt.Fscan(in, &t)
	s := []byte(t)
	ans := 0
	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			s[i] = opt(s[i-1], s[min(n-1, i+1)])
			ans++
		}
	}
	fmt.Fprintln(out, ans)
	fmt.Fprintln(out, string(s))
}

func opt(a, b byte) byte {
	if a == b {
		switch a {
		case 'R':

			return 'G'
		case 'G':
			return 'B'
		case 'B':
			return 'R'
		}
	} else {
		tmp := map[byte]bool{
			'R': false,
			'G': false,
			'B': false,
		}
		tmp[a] = true
		tmp[b] = true
		for c := range tmp {
			if !tmp[c] {
				return c
			}
		}
	}

	return 'R'
}
