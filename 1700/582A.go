package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	m := n * n
	a := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a[i])
	}

	if n == 1 {
		fmt.Fprintln(out, a[0])
		return
	}
	sort.Ints(a)
	ans := []int{a[m-1], a[m-2]}
	del := map[int]int{gcd(ans[0], ans[1]): 2}
	for i := m - 3; i >= 0; i-- {
		if del[a[i]] > 0 {
			del[a[i]]--
		} else {
			for _, x := range ans {
				del[gcd(x, a[i])] += 2
			}
			ans = append(ans, a[i])
		}
	}

	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
