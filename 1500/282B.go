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

	var a, g int
	sum := 0
	for range n {
		fmt.Fscan(in, &a, &g)
		if sum+a <= 500 {
			sum += a
			fmt.Fprint(out, "A")
		} else {
			sum -= g
			fmt.Fprint(out, "G")
		}
	}
}
