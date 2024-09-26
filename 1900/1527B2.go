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
		var s string
		fmt.Fscan(in, &s)
		cnt := 0
		for i := range s {
			if s[i] == '0' {
				cnt++
			}
		}
		if n&1 == 1 && s[n/2] == '0' && cnt == 2 {
			fmt.Fprintln(out, "DRAW")
			continue
		}
		p := 0
		for i := 0; i < n/2; i++ {
			if s[i] != s[n-i-1] {
				p++
			}
		}

		if p == 0 && !(n&1 == 1 && s[n/2] == '0' && cnt > 1) {
			fmt.Fprintln(out, "BOB")
		} else {
			fmt.Fprintln(out, "ALICE")
		}

	}

}
