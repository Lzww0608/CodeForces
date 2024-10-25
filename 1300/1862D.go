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

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	x := int(math.Sqrt(float64(n * 2)))
	for x*(x-1) <= n*2 {
		x++
	}
	x--
	fmt.Fprintln(out, x+n-(x-1)*x/2)
}
