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

	var n int
	fmt.Fscan(in, &n)
next:
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		x := a * b
		for j := int(math.Cbrt(float64(x))) - 1; ; j++ {
			if j*j*j == x {
				if a%j == 0 && b%j == 0 {
					fmt.Fprintln(out, "YES")
				} else {
					fmt.Fprintln(out, "NO")
				}
				continue next
			} else if j*j*j > x {
				break
			}
		}
		fmt.Fprintln(out, "NO")
	}
}
