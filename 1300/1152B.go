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

	var x int
	fmt.Fscan(in, &x)
	a := []int{}
	ans := 0
	for t := 0; (x+1)&x != 0; t++ {
		ans++
		if t&1 == 1 {
			x += 1
		} else {
			n := 0
			for x&(1<<n) == 0 {
				n++
			}
			x ^= (1 << n) - 1
			a = append(a, n)
		}

	}
	fmt.Fprintln(out, ans)
	for _, x := range a {
		fmt.Fprint(out, x, " ")
	}
}
