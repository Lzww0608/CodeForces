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
		cnt := 0
		for n > 0 {
			cnt += 1
			n -= cnt
		}
		if n == 0 || n < -1 {
			fmt.Fprintln(out, cnt)
			continue
		}
		fmt.Fprintln(out, cnt+1)

	}
}
