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

func calc(a int) []int {
	res := []int{}
	for i := 1; i*i <= a; i++ {
		if a%i == 0 {
			res = append(res, i)
			if i*i != a {
				res = append(res, a/i)
			}
		}

	}

	return res
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var a, b, c, d int
	fmt.Fscan(in, &a, &b, &c, &d)
	nums1 := calc(a)
	nums2 := calc(b)
	for _, i := range nums1 {
		for _, j := range nums2 {
			x := i * j
			y := a * b / x
			x = (a/x + 1) * x
			y = (b/y + 1) * y
			if x > a && x <= c && y > b && y <= d {
				fmt.Fprintln(out, x, y)
				return
			}
		}
	}
	fmt.Fprintln(out, -1, -1)
}
