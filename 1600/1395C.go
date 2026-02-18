package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 9

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, n)
	b := make([]int, m)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	for i := range b {
		fmt.Fscan(in, &b[i])
	}

next:
	for s := 0; s < (1 << N); s++ {
		for _, x := range a {
			f := false
			for _, y := range b {
				if (x&y)|s == s {
					f = true
					break
				}
			}
			if !f {
				continue next
			}
		}

		fmt.Fprintln(out, s)
		return
	}

	return
}
