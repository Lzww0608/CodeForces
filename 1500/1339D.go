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
	var s string
	fmt.Fscan(in, &s)

	ans := make([]int, n)
	pre0 := map[int]int{}
	pre1 := map[int]int{}
	k := 1
	for i := range s {
		if s[i] == '0' {
			if len(pre1) == 0 {
				pre0[i] = k
				ans[i] = k
				k++
			} else {
				for k, v := range pre1 {
					pre0[i] = v
					ans[i] = v
					delete(pre1, k)
					break
				}
			}
		} else {
			if len(pre0) == 0 {
				pre1[i] = k
				ans[i] = k
				k++
			} else {
				for k, v := range pre0 {
					pre1[i] = v
					ans[i] = v
					delete(pre0, k)
					break
				}
			}
		}
	}

	fmt.Fprintln(out, k-1)
	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)
}
