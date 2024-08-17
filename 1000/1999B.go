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

	var t, a, b, c, d int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &a, &b, &c, &d)
		ans := 0
		for i := 0; i < 2; i++ {
			a, b = b, a
			cnt := 0
			if a > c {
				cnt++
			} else if a < c {
				cnt--
			}

			if b > d {
				cnt++
			} else if b < d {
				cnt--
			}
			if cnt > 0 {
				ans += 2
			}
		}
		fmt.Fprintln(out, ans)
	}
}
