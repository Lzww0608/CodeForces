package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	//in := bufio.NewScanner(os.Stdin)

	var t, a, b int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &a, &b)
		Println(min(a, b, (a+b)/4))
	}

	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
