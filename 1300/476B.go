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

	var s, t string
	fmt.Fscan(in, &s, &t)
	a, b := 0, 0
	for _, c := range s {
		if c == '+' {
			a++
		} else {
			b++
		}
	}

	for _, c := range t {
		if c == '+' {
			a--
		} else if c == '-' {
			b--
		}
	}

	if a < 0 || b < 0 {
		fmt.Fprintf(out, "%f\n", 0.0)
		return
	}
	//fmt.Fprintln(out, a, b)
	fmt.Fprintf(out, "%.12f\n", float64(C(a, a+b))/float64(calc(a+b)))
}

func calc(a int) int {
	res := 1
	for i := 1; i <= a; i++ {
		res *= 2
	}

	return res
}

func C(a, b int) int {
	return factor(b) / (factor(a) * factor(b-a))
}

func factor(x int) int {
	if x == 0 {
		return 1
	}
	ans := 1
	for i := 1; i <= x; i++ {
		ans *= i
	}

	return ans
}
