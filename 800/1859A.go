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

	var t, n int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		a := make([]int, n)
		mn, cnt := math.MaxInt32, 0
		for i := 0; i < n; i++ {
			Fscan(in, &a[i])
			if a[i] < mn {
				mn, cnt = a[i], 1
			} else if a[i] == mn {
				cnt++
			}
		}
		if cnt == n {
			Println(-1)
			continue
		}
		Println(cnt, n-cnt)
		for i := 0; i < cnt; i++ {
			Print(mn, " ")
		}
		Println()
		for _, x := range a {
			if x != mn {
				Print(x, " ")
			}

		}
		Println()

	}

	return
}
