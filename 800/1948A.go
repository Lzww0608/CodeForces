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
			Println("NO")
			continue
		}
		ans := []byte{}
		for i := 0; i < n; i++ {
			ans = append(ans, byte('A'+(i/2)%26))
		}

		Println("YES")
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
