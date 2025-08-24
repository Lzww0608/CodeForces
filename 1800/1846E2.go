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
	for i := 2; i <= 60; i++ {
		k := 1
		if i == 2 {
			k = int(math.Sqrt(float64(n)))
		} else if i == 3 {
			k = int(math.Cbrt(float64(n)))
		} else {
			k = int(math.Pow(float64(n), 1.0/float64(i)))
		}

		if k == 1 {
			continue
		}

		sum, p := 1, 1
		for t := 1; t <= i; t++ {
			p *= k
			sum += p
		}

		if sum == n {
			fmt.Fprintln(out, "YES")
			return
		}
	}
	fmt.Fprintln(out, "NO")
}
