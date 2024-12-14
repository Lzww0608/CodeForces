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

	var n int
	fmt.Fscan(in, &n)

	deg := make([]int, n)
	xor := make([]int, n)
	used := 0
	q := []int{}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &deg[i], &xor[i])
		if deg[i] == 1 {
			q = append(q, i)
		} else if deg[i] == 0 {
			if xor[i] != 0 {
				panic("false")
			}
			used++
		}
	}

	ans := [][2]int{}
	for len(q) > 0 {
		from := q[0]
		q = q[1:]
		used++
		if deg[from] == 0 {
			if xor[from] != 0 {
				panic("false")
			}
			continue
		}
		deg[from]--
		to := xor[from]
		xor[from] = 0
		if to >= n {
			panic("false")
		}
		ans = append(ans, [2]int{from, to})
		xor[to] ^= from
		if deg[to] == 0 {
			panic("false")
		}
		deg[to]--
		if deg[to] == 1 {
			q = append(q, to)
		}
	}

	if used != n {
		panic("false")
	}

	fmt.Fprintln(out, len(ans))
	for _, v := range ans {
		fmt.Fprintln(out, v[0], v[1])
	}
}
