package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	//in := bufio.NewScanner(os.Stdin)

	var t, n, x int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		mn, cnt := math.MaxInt32, 0
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			if x < mn {
				mn, cnt = x, 1
			} else if x == mn {
				cnt++
			}
		}
		Println(n - cnt)
	}

	return
}
