package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	//in := bufio.NewScanner(os.Stdin)

	var t, n, m, k, h, x int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &m, &k, &h)
		ans := 0
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			for j := 1; j < m; j++ {
				if abs(x-h) == j*k {
					ans++
					break
				}
			}
		}
		Println(ans)
	}

	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
