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

	var t, n int

	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		s := make([]string, 2)
		for i := range s {
			fmt.Fscan(in, &s[i])
		}
		ans := 0
		for i := 1; i < n-1; i++ {
			for j := 0; j < 2; j++ {
				if s[j][i] == '.' && s[j][i-1] == '.' && s[j][i+1] == '.' && s[j^1][i] == '.' && s[j^1][i+1] == 'x' && s[j^1][i-1] == 'x' {
					ans++
				}
			}
		}

		fmt.Fprintln(out, ans)
	}
}
