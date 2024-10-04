package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t, n int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		a := make([]int, 2*n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}
		sort.Ints(a)
		l, r := 0, len(a)-1
		for i := range a {
			if i&1 == 0 {
				fmt.Print(a[l], " ")
				l++
			} else {
				fmt.Print(a[r], " ")
				r--
			}
		}
	}
	return
}
