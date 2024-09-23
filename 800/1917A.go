package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t, n int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		a := make([]int, n)
		pos, idx := 0, 1
		zero := false
		for i := range a {
			fmt.Fscan(in, &a[i])
			if a[i] < 0 {
				idx = i + 1
				pos++
			} else if a[i] == 0 {
				zero = true
			}
		}
		if pos&1 == 1 || zero {
			fmt.Println(0)
		} else {
			fmt.Println(1)
			fmt.Println(idx, 0)
		}
	}
}
