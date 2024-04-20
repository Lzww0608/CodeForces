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
		if n&1 == 0 {
			Println(0)
			continue
		}
		f := false
		for n >= 10 {
			n /= 10
			if n&1 == 0 {
				f = true
			}
		}
		if n&1 == 0 {
			Println(1)
		} else if f {
			Println(2)
		} else {
			Println(-1)
		}
	}

	return
}
