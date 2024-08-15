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

		ans, cnt := 0, 0
		for _, c := range s {
			if c == '1' {
				cnt++
			} else if cnt > 0 {
				ans += cnt + 1
			}
		}
		fmt.Fprintln(out, ans)
	}
}
