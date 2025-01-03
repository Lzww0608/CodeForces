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

	var n, x, y, z int
	fmt.Fscan(in, &n)
	a := []int{0}
	b := []int{0}
	sum := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		if x == 1 {
			fmt.Fscan(in, &y, &z)
			b[y-1] += z
			sum += z * y
		} else if x == 2 {
			fmt.Fscan(in, &y)
			a = append(a, y)
			b = append(b, 0)
			sum += y
		} else {
			sum -= a[len(a)-1] + b[len(b)-1]
			b[len(b)-2] += b[len(b)-1]
			a = a[:len(a)-1]
			b = b[:len(b)-1]
		}

		fmt.Fprintf(out, "%.6f\n", float64(sum)/float64(len(a)))
	}

}
