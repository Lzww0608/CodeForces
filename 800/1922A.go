package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	//in := bufio.NewScanner(os.Stdin)
	/*
		var s string
		in.Scan()
		s = in.Text()
	*/

	var n, t int
next:
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		var a, b, c string
		Fscan(in, &a, &b, &c)
		for i := range a {
			if c[i] != b[i] && c[i] != a[i] {
				Println("YES")
				continue next
			}
		}
		Println("NO")
	}

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
