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
		ans := []byte{}
		var s string
		Fscan(in, &s)
		for i := 0; i < 2*n-1; i++ {
			if i&1 == 0 {
				ans = append(ans, s[i])
			}
		}
		Println(string(ans))
	}

	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
