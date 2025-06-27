package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	m := make(map[string]int)
	mx := math.MinInt32

	s := make([]string, n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i], &a[i])
		m[s[i]] += a[i]

	}

	for _, v := range m {
		mx = max(mx, v)
	}

	mm := make(map[string]int)
	for i := 0; i < n; i++ {
		mm[s[i]] += a[i]
		if mm[s[i]] >= mx && m[s[i]] == mx {
			fmt.Fprintln(out, s[i])
			return
		}
	}

}
