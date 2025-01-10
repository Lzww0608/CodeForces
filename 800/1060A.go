package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 10

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	var s string
	fmt.Fscan(in, &s)
	cnt := make([]int, N)
	for _, c := range s {
		cnt[c-'0']++
	}

	ans := 0
	for i := cnt[8]; i > 0; i-- {
		ans = max(ans, min(i, (n-i)/10))
	}
	fmt.Fprintln(out, ans)
}
