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

	var n, x, y int
	fmt.Fscan(in, &n)
	v := make([][]int, n)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(in, &x, &y)
		x--
		y--
		v[x] = append(v[x], i)
		v[y] = append(v[y], i)
	}

	mx := []int{0, 0}
	for i := range v {
		if len(v[i]) > mx[0] {
			mx = []int{len(v[i]), i}
		}
	}
	cur := 0
	ans := make([]int, n-1)
	for i := range ans {
		ans[i] = -1
	}
	for _, t := range v[mx[1]] {
		ans[t] = cur
		cur++
	}
	for i := range ans {
		if ans[i] == -1 {
			ans[i] = cur
			cur++
		}
	}
	for _, x := range ans {
		fmt.Fprintln(out, x)
	}
}
