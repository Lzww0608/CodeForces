package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	//in := bufio.NewScanner(os.Stdin)

	var t, n int64
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		h1, h2, h3 := n / 3, n / 3, n / 3
		if n % 3 == 0 {
			h1++
			h3--
		} else if n % 3 == 1 {
			h1 += 2
			h3--
		} else {
			h1 += 2
			h2++
			h3--
		}
		Println(h2, h1, h3)
	}

	return
}
