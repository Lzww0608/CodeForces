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

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, 3)

	ans := make([]int, n)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a[0], &a[1], &a[2])
		a[0]--
		a[1]--
		a[2]--
		cnt := ans[a[0]] | ans[a[1]] | ans[a[2]]
		k := 1
		for _, x := range a {
			if k == cnt {
				k++
			}
			if ans[x] == 0 {
				ans[x] = k
				k++
			}
		}
	}

	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
}
