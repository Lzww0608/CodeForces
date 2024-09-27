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

	var t, n int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		b := make([]int, n+2)
		for i := range b {
			fmt.Fscan(in, &b[i])
		}
		sort.Ints(b)
		sum := 0
		for i := 0; i < n; i++ {
			sum += b[i]
		}
		if sum == b[n] || sum == b[n+1] {
			for i := 0; i < n; i++ {
				fmt.Fprint(out, b[i], " ")
			}
			fmt.Fprintln(out)
			continue
		}
		sum += b[n]
		idx := -1
		for i := 0; i < n; i++ {
			if sum-b[i] == b[n+1] {
				idx = i
				break
			}
		}
		if idx == -1 {
			fmt.Fprintln(out, -1)
			continue
		}
		for i := 0; i < n+1; i++ {
			if i != idx {
				fmt.Fprint(out, b[i], " ")
			}
		}
		fmt.Fprintln(out)
	}
}
