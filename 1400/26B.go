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

	var s string
	fmt.Fscan(in, &s)
	ans, cnt := 0, 0
	for _, c := range s {
		if c == '(' {
			cnt++
		} else if c == ')' && cnt > 0 {
			ans += 2
			cnt--
		}
	}
	fmt.Fprintln(out, ans)
}
