package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 26

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
	var s string
	fmt.Fscan(in, &s)
	n := len(s)
	cnt := 0

	f := [N]bool{}
	for i := 0; i < n; i++ {
		x := int(s[i] - 'a')
		if f[x] {
			cnt += 2
			for j := range f {
				f[j] = false
			}
		} else {
			f[x] = true
		}
	}

	fmt.Fprintln(out, n-cnt)
}
