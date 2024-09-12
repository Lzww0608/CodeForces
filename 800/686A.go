package main

import (
	. "fmt"
)

func main() {

	//in := bufio.NewReader(os.Stdin)
	//in := bufio.NewScanner(os.Stdin)
	var n, x, ans, cnt int64
	var c rune

	Scanf("%d %d\n", &n, &x)
	ans, cnt = x, 0
	for i := int64(0); i < n; i++ {
		Scanf("%c %d\n", &c, &x)
		if c == '+' {
			ans += x
		} else if ans >= x {
			ans -= x
		} else {
			cnt++
		}

	}

	Println(ans, cnt)

	return
}
