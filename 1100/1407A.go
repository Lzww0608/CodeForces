package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t, n int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		a := make([]int, n)
		zero, one := 0, 0
		for i := range a {
			fmt.Fscan(in, &a[i])
			if a[i] == 0 {
				zero++
			} else {
				one++
			}
		}

		if zero >= one {
			fmt.Fprintln(out, zero)
			fmt.Fprintln(out, strings.Repeat("0 ", zero))
		} else {
			one &^= 1
			fmt.Fprintln(out, one)
			fmt.Fprintln(out, strings.Repeat("1 ", one))
		}

	}
}
