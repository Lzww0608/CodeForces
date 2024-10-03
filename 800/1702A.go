package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	//in := bufio.NewScanner(os.Stdin)

	var t, m int64
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &m)
		var k int64 = 1e9
		for k > m {
			k /= 10
		}
		Println(m - k)
	}

	return
}
