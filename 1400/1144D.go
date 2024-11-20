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

	//var t int
	//for fmt.Fscan(in, &t); t > 0; t-- {
	solve(in, out)
	//}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	m := map[int]int{}
	for i := range a {
		fmt.Fscan(in, &a[i])
		m[a[i]]++
	}
	ans := [][]int{}
	mx := 0
	for k, v := range m {
		if v > m[mx] {
			mx = k
		}
	}
	pos := -1
	for i := range a {
		if a[i] == mx {
			pos = i
			break
		}
	}

	l, r := pos-1, pos+1
	for l >= 0 {
		x := a[l]
		if x < mx {
			ans = append(ans, []int{1, l + 1, l + 2})
		} else if x > mx {
			ans = append(ans, []int{2, l + 1, l + 2})
		}
		l--
	}

	for r < n {
		x := a[r]
		if x < mx {
			ans = append(ans, []int{1, r + 1, r})
		} else if x > mx {
			ans = append(ans, []int{2, r + 1, r})
		}
		r++
	}

	fmt.Fprintln(out, len(ans))
	for _, v := range ans {
		fmt.Fprintln(out, v[0], v[1], v[2])
	}
}
