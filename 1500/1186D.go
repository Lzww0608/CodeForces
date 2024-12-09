package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const EPS float64 = 1e-6

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, sum int
	fmt.Fscan(in, &n)
	a := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		sum += int(a[i])
	}

	for i := range a {
		if sum > 0 && a[i] < 0 && math.Abs(a[i]-float64(int(a[i]))) > EPS {
			fmt.Fprintln(out, int(a[i])-1)
			sum--
		} else if sum < 0 && a[i] > 0 && math.Abs(a[i]-float64(int(a[i]))) > EPS {
			fmt.Fprintln(out, int(a[i])+1)
			sum++
		} else {
			fmt.Fprintln(out, int(a[i]))
		}
	}
}
