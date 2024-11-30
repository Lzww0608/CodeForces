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
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	m := map[int]bool{}
	sum, ans := 0, 0
	m[0] = true
	for _, x := range a {
		sum += x
		if m[sum] {
			ans++
			sum = x
			clear(m)
			m[0] = true
		}
		m[sum] = true
	}

	fmt.Fprintln(out, ans)
}
