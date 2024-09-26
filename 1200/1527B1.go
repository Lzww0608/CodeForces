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
		if cnt == 1 || cnt&1 == 0 {
			fmt.Fprintln(out, "BOB")
		} else {
			fmt.Fprintln(out, "ALICE")
		}

	}
}
