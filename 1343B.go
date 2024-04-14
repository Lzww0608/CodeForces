package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	//in := bufio.NewScanner(os.Stdin)
	var t, n int

	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		if n%4 != 0 {
			Println("NO")
			continue
		}
		Println("YES")
		for i, x := 1, 2; i <= n/2; i++ {
			Print(x, " ")
			if i&1 == 0 {
				x += 4
			} else {
				x += 2
			}
		}

		for i, x := 1, 1; i <= n/2; i++ {
			Print(x, " ")
			if i&1 == 0 {
				x += 2
			} else {
				x += 4
			}
		}
		Println()
	}

	return
}
