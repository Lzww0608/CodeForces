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

	var n, m, k int
	fmt.Fscan(in, &n, &m)
	var str, c string
	fmt.Fscan(in, &str)
	s := append([]byte{0}, append([]byte(str), 0)...)
	dot, seg := 0, 0
	for i := range s {
		if s[i] == '.' {
			dot++
			if s[i-1] != '.' {
				seg++
			}
		}
	}

	for i := 0; i < m; i++ {
		fmt.Fscan(in, &k, &c)
		if (s[k] == '.') != (c[0] == '.') {
			d := -1
			if s[k] != '.' {
				d = 1
			}
			dot += d
			if s[k-1] == '.' && s[k+1] == '.' {
				seg -= d
			} else if s[k-1] != '.' && s[k+1] != '.' {
				seg += d
			}
			s[k] = c[0]
		}
		fmt.Fprintln(out, dot-seg)
	}
}
