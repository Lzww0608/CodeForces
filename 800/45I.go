package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
)

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	Fscan(in, &n)
	a := make([]int, n)
	id, neg, mx := 0, 0, math.MinInt32
	zero := 0
	for i := range a {
		Fscan(in, &a[i])
		if a[i] < 0 {
			neg++
			if a[i] > mx {
				mx = a[i]
				id = i
			}
		} else if a[i] == 0 {
			zero++
		}
	}

	if n == 1 {
		Print(a[0])
	} else if zero == n || (zero == n-1 && neg == 1) {
		Print(0)
	} else {
		for i, x := range a {
			if x != 0 && (neg%2 == 0 || id != i) {
				Print(x, " ")
			}
		}
	}

	return
}
