package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k int
	fmt.Fscan(in, &n, &k)
	var s string
	fmt.Fscan(in, &s)

	ans := strings.Repeat(s[:1], k)
	for d := 2; d <= min(n, k); d++ {
		t := (k + d - 1) / d
		cur := strings.Repeat(s[:d], t)
		ans = min(ans, cur[:k])
	}
	fmt.Fprintln(out, ans)
}
