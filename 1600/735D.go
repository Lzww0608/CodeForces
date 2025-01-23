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

	var n int
	fmt.Fscan(in, &n)
	if isPrime(n) {
		fmt.Fprintln(out, 1)
	} else if n&1 == 0 {
		fmt.Fprintln(out, 2)
	} else {
		if isPrime(n - 2) {
			fmt.Fprintln(out, 2)
		} else {
			fmt.Fprintln(out, 3)
		}
	}
}
func isPrime(n int) bool {
	if n == 2 {
		return true
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
