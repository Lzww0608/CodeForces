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
	var r int
	fmt.Fscan(in, &r)
	ans := 0
	mx, mn := (r+1)*(r+1), r*r

	j := r
	for i := range r + 1 {
		for i*i+j*j >= mx {
			j--
		}

		cur := j
		//fmt.Fprintln(out, j)
		for i*i+cur*cur >= mn && cur > 0 {
			ans++
			cur--
		}
	}

	fmt.Fprintln(out, ans*4)
}
