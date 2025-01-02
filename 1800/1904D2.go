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
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	b := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	for i := range b {
		fmt.Fscan(in, &b[i])
	}

	for i := 0; i < n; i++ {
		if a[i] > b[i] {
			fmt.Fprintln(out, "NO")
			return
		}
	}

	l := make([]int, n)
	r := make([]int, n)
	st := []int{}
	for i, x := range a {
		for len(st) > 0 && a[st[len(st)-1]] <= x {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			l[i] = st[len(st)-1]
		} else {
			l[i] = -1
		}
		st = append(st, i)
	}

	st = []int{}
	for i := n - 1; i >= 0; i-- {
		x := a[i]
		for len(st) > 0 && a[st[len(st)-1]] <= x {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			r[i] = st[len(st)-1]
		} else {
			r[i] = n
		}
		st = append(st, i)
	}

	j := 0
	for i := 0; i < n; i++ {
		for j < n && a[i] == b[j] && j > l[i] && j < r[i] {
			j++
		}
	}

	if j == n {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}

}
