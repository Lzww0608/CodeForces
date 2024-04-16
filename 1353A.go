package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	//in := bufio.NewScanner(os.Stdin)
	var t, n, m int64

	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &m)
		var ans int64 = 0
		if n == 2 {
			ans += m
		} else if n > 2 {
			ans += m * 2
		}

		Println(ans)
	}

	return
}
