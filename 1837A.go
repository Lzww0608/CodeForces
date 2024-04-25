package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	//in := bufio.NewScanner(os.Stdin)

	var t, x, k int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &x, &k)
		if x%k != 0 {
			Println(1)
			Println(x)
		} else {
			Println(2)
			Println(1, x-1)
		}

	}

	return
}
