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

	var n, x int
	fmt.Fscan(in, &n)

	ans := 0
	cnt := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		cnt[x-i] += x
		ans = max(ans, cnt[x-i])
	}

	fmt.Fprintln(out, ans)
}
