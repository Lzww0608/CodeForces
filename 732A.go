package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	//in := bufio.NewScanner(os.Stdin)
	var k, r int
	Fscan(in, &k, &r)
	ans := k
	for i := 1; i <= 10; i++ {
		ans = k * i
		if ans%10 == 0 || (ans-r)%10 == 0 {
			break
		}
	}

	Println(ans / k)
	return
}
