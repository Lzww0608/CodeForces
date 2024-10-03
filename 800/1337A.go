package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	//in := bufio.NewScanner(os.Stdin)
	var t, a, b, c, d int64

	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &a, &b, &c, &d)
		Println(b, c, c)

	}

	return
}
