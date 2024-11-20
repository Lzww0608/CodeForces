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

	var s string
	fmt.Fscan(in, &s)
	a, b := 0, 0
	for _, c := range s {
		if c == '0' {
			fmt.Fprintln(out, 1, a%4+1)
			a++
		} else {
			fmt.Fprintln(out, 3, b%2*2+1)
			b++
		}
	}
}
