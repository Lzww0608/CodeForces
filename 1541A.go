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
		if n&1 == 1 {
			Print(3, 1, 2, " ")
			for i := 4; i <= n; i++ {
				if i&1 == 0 {
					Print(i+1, " ")
				} else {
					Print(i-1, " ")
				}
			}
		} else {
			for i := 1; i <= n; i++ {
				if i&1 == 1 {
					Print(i+1, " ")
				} else {
					Print(i-1, " ")
				}
			}
		}

		Println()

	}

	return
}
