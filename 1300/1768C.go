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
	p := make([]int, n)
	q := make([]int, n)
	cnt := make([]int, n)
	pos := make([][]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
		a[i]--
		p[i] = a[i]
		q[i] = a[i]
		cnt[a[i]]++
		pos[a[i]] = append(pos[a[i]], i)
	}
	zero := []int{}
	two := []int{}
	for i := 0; i < n; i++ {
		if cnt[i] == 0 {
			zero = append(zero, i)
		} else if cnt[i] == 2 {
			two = append(two, i)
		} else if cnt[i] > 2 {
			fmt.Fprintln(out, "NO")
			return
		}
	}

	for i := 0; i < len(two); i++ {
		if two[i] < zero[i] {
			fmt.Fprintln(out, "NO")
			return
		}
		v := two[i]
		p[pos[v][0]] = zero[i]
		q[pos[v][1]] = zero[i]
	}
	fmt.Fprintln(out, "YES")
	for i := 0; i < n; i++ {
		fmt.Fprint(out, p[i]+1, " ")
	}
	fmt.Fprintln(out)
	for i := 0; i < n; i++ {
		fmt.Fprint(out, q[i]+1, " ")
	}
	fmt.Fprintln(out)

}
