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

	var k int
	fmt.Fscan(in, &k)
	if k&1 == 1 {
		fmt.Fprintln(out, -1)
		return
	}

	a, b := "b", "w"
	for i := 0; i < k; i++ {
		for j := 0; j < k; j++ {
			for z := 0; z < k; z++ {
				if (j/2+z/2)&1 == 1 {
					fmt.Fprint(out, a)
				} else {
					fmt.Fprint(out, b)
				}
			}
			fmt.Fprintln(out)
		}
		a, b = b, a
		fmt.Fprintln(out)
	}
}
