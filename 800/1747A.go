package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t, n int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		var x, ans int = 0, 0
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			ans += x
		}
		Println(abs(ans))
	}

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
