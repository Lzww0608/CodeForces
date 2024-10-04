package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, t int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		for _, x := range a {
			if x&1 == 0 {
				Print(x, " ")
			}
		}
		for _, x := range a {
			if x&1 == 1 {
				Print(x, " ")
			}
		}
		Println()
	}

}
