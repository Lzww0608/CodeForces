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

	var n, s, u, v int
	fmt.Fscan(in, &n, &s)
	deg := make([]int, n)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(in, &u, &v)
		u--
		v--
		deg[u]++
		deg[v]++
	}

	cnt := 0
	for _, x := range deg {
		if x == 1 {
			cnt++
		}
	}

	fmt.Fprintln(out, float64(s*2)/float64(cnt))
}
