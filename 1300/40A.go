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

	var a, b int
	fmt.Fscan(in, &a, &b)
	d := math.Sqrt(float64(a*a + b*b))
	if d-float64(int(d)) <= EPS {
		fmt.Fprintln(out, "black")
		return
	}
	if a*b >= 0 {
		if int(d)&1 == 1 {
			fmt.Fprintln(out, "white")
		} else {
			fmt.Fprintln(out, "black")
		}
	} else {
		if int(d)&1 == 1 {
			fmt.Fprintln(out, "black")
		} else {
			fmt.Fprintln(out, "white")
		}
	}

}
