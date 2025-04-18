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

	var n, h int
	fmt.Fscan(in, &h, &n)
	n--
	dir := 0
	ans := 0
	for h > 0 {
		h--
		if ((n>>h)&1)^dir == 1 {
			ans += 1 << (h + 1)
		} else {
			ans++
			dir ^= 1
		}
	}

	fmt.Fprintln(out, ans)
}
