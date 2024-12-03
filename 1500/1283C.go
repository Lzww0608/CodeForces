package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	vis := make([]bool, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		a[i]--
		if a[i] >= 0 {
			vis[a[i]] = true
		}
	}
	b := make([]int, 0, n)
	c := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if a[i] == -1 {
			b = append(b, i)
		}
		if !vis[i] {
			c = append(c, i)
		}
	}
	//fmt.Fprintln(out, b, c)
	for {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(b), func(i, j int) {
			b[i], b[j] = b[j], b[i]
		})
		ok := true
		for i := 0; i < len(b); i++ {
			a[b[i]] = c[i]
			if b[i] == c[i] {
				ok = false
				break
			}
		}
		if ok {
			break
		}
	}
	for _, x := range a {
		fmt.Fprint(out, x+1, " ")
	}

}
