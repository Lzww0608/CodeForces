package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 200_001

var f [N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	f[1] = 1
	pre := 3
	for i := 2; i < N; i++ {
		if i == pre {
			pre *= 3
			f[i] = f[i-1] + 1
		} else {
			f[i] = f[i-1]
		}
	}

	for i := 1; i < N; i++ {
		f[i] += f[i-1]
	}

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var l, r int
	fmt.Fscan(in, &l, &r)

	fmt.Fprintln(out, f[r]-f[l-1]+f[l]-f[l-1])
}
