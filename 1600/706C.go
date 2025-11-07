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
	c := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &c[i])
	}

	s := make([]string, n)
	t := make([][]byte, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
		for j := len(s[i]) - 1; j >= 0; j-- {
			t[i] = append(t[i], s[i][j])
		}
	}

	f := make([][2]int, n)
	for i := range n {
		f[i] = [2]int{math.MaxInt / 2, math.MaxInt / 2}
	}
	f[0] = [2]int{0, c[0]}

	for i := 1; i < n; i++ {
		if string(s[i]) >= string(s[i-1]) {
			f[i][0] = min(f[i][0], f[i-1][0])
		}
		if string(s[i]) >= string(t[i-1]) {
			f[i][0] = min(f[i][0], f[i-1][1])
		}
		if string(t[i]) >= string(s[i-1]) {
			f[i][1] = min(f[i][1], f[i-1][0]+c[i])
		}
		if string(t[i]) >= string(t[i-1]) {
			f[i][1] = min(f[i][1], f[i-1][1]+c[i])
		}

		if f[i][0] == math.MaxInt / 2 && f[i][1] == math.MaxInt / 2 {
			fmt.Fprintln(out, -1)
			return
		}
	}

	fmt.Fprintln(out, min(f[n-1][0], f[n-1][1]))
}
