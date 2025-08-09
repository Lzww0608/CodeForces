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
	xor := 0
	for i := range a {
		fmt.Fscan(in, &a[i])
		xor ^= a[i]
	}
	if xor == 0 {
		fmt.Fprintln(out, "YES")
		return
	}

	cnt := 0
	for i := 0; i < n; i++ {
		for s := 0; i < n && s^a[i] != xor; i++ {
			s ^= a[i]
		}
		cnt++
	}

	if cnt > 2 {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}
}
