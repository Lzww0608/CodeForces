package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	//in := bufio.NewScanner(os.Stdin)
	var t, l, r int64

	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &l, &r)
		if r < l*2 {
			Println(-1, -1)
		} else {
			Println(l, l*2)
		}

	}

	return
}
