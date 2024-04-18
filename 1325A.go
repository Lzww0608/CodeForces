package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	//in := bufio.NewScanner(os.Stdin)
	var t, x int64

	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &x)
		Println(1, x-1)

	}

	return
}
