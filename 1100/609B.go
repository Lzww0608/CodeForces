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

	var n, m, x int
	fmt.Fscan(in, &n, &m)
	cnt := make([]int, m)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		cnt[x-1]++
	}
	ans := 0
	for _, x := range cnt {
		ans += (n - x) * x
	}
	fmt.Fprintln(out, ans/2)
}
