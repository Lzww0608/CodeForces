package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {

	var s string
	//in := bufio.NewReader(os.Stdin)
	in := bufio.NewScanner(os.Stdin)
	if in.Scan() {
		s = in.Text()
	}
	ans := 0
	m := map[byte]bool{}
	for i := 1; i < len(s)-1; i += 3 {
		if m[s[i]] == false {
			m[s[i]] = true
			ans++
		}
	}
	Println(ans)
	return
}
