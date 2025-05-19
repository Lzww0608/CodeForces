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
	sum := 0
	for i := range a {
		fmt.Fscan(in, &a[i])
		sum ^= a[i]
	}

	if sum == 0 {
		fmt.Fprintln(out, "DRAW")
		return
	}

	for k := 30; k >= 0; k-- {
		f := [2]int{}
		for _, x := range a {
			f[(x>>k)&1]++
		}
		if f[1]%2 == 0 {
			continue
		}
		if f[1]%4 == 3 && f[0]%2 == 0 {
			fmt.Fprintln(out, "LOSE")
		} else {
			fmt.Fprintln(out, "WIN")
		}
		return
	}
}
