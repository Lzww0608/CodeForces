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

	var m, t, x int
	cnt := make([]int, 30)
	fmt.Fscan(in, &m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &t, &x)
		if t == 1 {
			cnt[x]++
		} else {
			cur := 0
			for i := 29; i >= 0; i-- {
				cur += (x >> i) & 1
				cur -= min(cur, cnt[i])
				cur *= 2
			}
			if cur == 0 {
				fmt.Fprintln(out, "YES")
			} else {
				fmt.Fprintln(out, "NO")
			}
		}
	}

}
