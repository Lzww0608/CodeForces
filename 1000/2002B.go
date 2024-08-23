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

	var t, n int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		a := make([]int, n)
		b := make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}
		for i := range b {
			fmt.Fscan(in, &b[i])
		}
		f := true
		for i := range a {
			if a[i] != b[i] {
				f = false
				break
			}
		}
		if f {
			fmt.Fprintln(out, "Bob")
			continue
		}
		f = true
		for i := range a {
			if a[i] != b[n-i-1] {
				f = false
				break
			}
		}
		if f {
			fmt.Fprintln(out, "Bob")
			continue
		}
		fmt.Fprintln(out, "Alice")
	}
}
