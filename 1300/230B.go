package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const N int = 1_000_001

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	isPrimes := make([]bool, N)
	isPrimes[1] = true
	for i := 2; i < N; i++ {
		if !isPrimes[i] {
			for j := i * i; j < N; j += i {
				isPrimes[j] = true
			}
		}
	}

	var n, x int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		y := int(math.Sqrt(float64(x)))
		if y*y == x && !isPrimes[y] {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}

}
