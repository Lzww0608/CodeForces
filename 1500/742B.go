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
	fmt.Fscan(in, &n, &x)
	a := make([]int, n)
	cnt := make(map[int]int)
	ans := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		ans += cnt[x^a[i]]
		cnt[a[i]]++
	}
	fmt.Fprintln(out, ans)
}
