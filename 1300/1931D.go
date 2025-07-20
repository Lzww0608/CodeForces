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
		solve(in, out)
	}

}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n, x, y int
	fmt.Fscan(in, &n, &x, &y)
	a := make([]int, n)
	type pair struct {
		x, y int 
	}
	cnt := make(map[pair]int)

	ans := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		ans += cnt[pair{a[i] % x, a[i] % y}]
		cnt[pair{(x - a[i] % x) % x, a[i] % y}]++
	}

	fmt.Fprintln(out, ans)
}