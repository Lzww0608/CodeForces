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
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])

		}
		if a[0] == a[n-1] {
			Println("NO")
			continue
		}
		Println("YES")
		Print(a[0], " ")
		for i := n - 1; i > 0; i-- {
			Print(a[i], " ")
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
