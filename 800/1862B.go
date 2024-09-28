package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	//in := bufio.NewScanner(os.Stdin)

	var t, n, x int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		a := []int{}
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			if len(a) == 0 || x >= a[len(a)-1] {
				a = append(a, x)
			} else {
				a = append(a, 1)
				a = append(a, x)
			}
		}
		Println(len(a))
		for _, t := range a {
			Print(t, " ")
		}
		Println()
	}

	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
