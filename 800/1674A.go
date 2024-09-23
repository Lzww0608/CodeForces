package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t, x, y int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &x, &y)
		if x > y || y%x != 0 {
			Println(0, 0)
			continue
		}
		Println(1, y/x)

	}
}
