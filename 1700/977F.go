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

	f := make(map[int]int)
	mx := 0
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		if f[a[i]] = f[a[i]-1] + 1; f[a[i]] > f[mx] {
			mx = a[i]
		}
	}

	d := mx - f[mx] + 1
	ans := []any{}
	for i, x := range a {
		if x == d {
			d++
			ans = append(ans, i+1)
		}
	}
	fmt.Fprintln(out, len(ans))
	fmt.Fprint(out, ans...)
}
