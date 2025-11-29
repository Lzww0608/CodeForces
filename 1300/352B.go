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
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	d := make(map[int]int)
	pos := make(map[int]int)
	for i, x := range a {
		if p, ok := pos[x]; ok {
			if y, okk := d[x]; okk {
				if y != i-p {
					d[x] = -1
				}
			} else {
				d[x] = i - p
			}
		}
		pos[x] = i
	}

	ans := []int{}
	for k := range pos {
		if v, ok := d[k]; !ok || v != -1 {
			ans = append(ans, k)
		}
	}
	sort.Ints(ans)
	fmt.Fprintln(out, len(ans))
	for _, v := range ans {
		fmt.Fprintln(out, v, d[v])
	}
}
