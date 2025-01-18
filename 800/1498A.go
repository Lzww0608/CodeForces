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

	var t, x int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &x)
		sum := digitSum(x)
		if gcd(x, sum) > 1 {
			fmt.Fprintln(out, x)
		} else if gcd(x+1, digitSum(x+1)) > 1 {
			fmt.Fprintln(out, x+1)
		} else {
			fmt.Fprintln(out, x+2)
		}

	}
}

func digitSum(x int) (ans int) {
	for x > 0 {
		ans += x % 10
		x /= 10
	}
	return
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
