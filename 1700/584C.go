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

	var n, t int
	fmt.Fscan(in, &n, &t)
	var s1, s2 string
	fmt.Fscan(in, &s1, &s2)
	same := 0
	for i := 0; i < n; i++ {
		if s1[i] == s2[i] {
			same++
		}
	}

	if (n - same) > t*2 {
		fmt.Fprintln(out, -1)
		return
	}

	if t >= n-same {
		k := t - (n - same)
		for i := 0; i < n; i++ {
			if s1[i] == s2[i] {
				if k > 0 {
					k--
					fmt.Fprintf(out, "%c", diff(s1[i], s2[i]))
				} else {
					fmt.Fprintf(out, "%c", s1[i])
				}
			} else {
				fmt.Fprintf(out, "%c", diff(s1[i], s2[i]))
			}
		}
	} else {
		a, b := n-same-t, n-same-t
		for i := 0; i < n; i++ {
			if s1[i] == s2[i] {
				fmt.Fprintf(out, "%c", s1[i])
			} else {
				if a > 0 {
					a--
					fmt.Fprintf(out, "%c", s1[i])
				} else if b > 0 {
					b--
					fmt.Fprintf(out, "%c", s2[i])
				} else {
					fmt.Fprintf(out, "%c", diff(s1[i], s2[i]))
				}
			}
		}
	}
}

func diff(a, b byte) byte {
	for x := byte('a'); x <= 'z'; x++ {
		if x != a && x != b {
			return x
		}
	}

	return 'a'
}
