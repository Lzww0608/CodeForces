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
	a := make([]int, n)
	b := make([]int, n)
	pos := []int{}
	for i := range a {
		b[i] = n
		fmt.Fscan(in, &a[i])
		if a[i] == 0 {
			pos = append(pos, i)
		}
	}
	for _, x := range pos {
		b[x] = 0
		l, r := x-1, x+1
		for k := 1; l >= 0; k++ {
			if b[l] >= k && a[l] != 0 {
				b[l] = k
			} else {
				break
			}
			l--
		}
		for k := 1; r < n; k++ {
			if b[r] >= k && a[r] != 0 {
				b[r] = k
			} else {
				break
			}
			r++
		}
	}
	for _, x := range b {
		fmt.Fprintln(out, x)
	}
}
