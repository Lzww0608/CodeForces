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
		var s string
		fmt.Fscan(in, &s)
		cnt := 1 - int(s[0]-'0')
		for i := 1; i < len(s); i++ {
			if s[i] != s[i-1] && s[i] == '0' {
				cnt++
			}
		}
		fmt.Fprintln(out, min(cnt, 2))
	}
}
