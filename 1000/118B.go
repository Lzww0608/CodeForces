package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	s := []string{}
	for i := 0; i <= n; i++ {
		tmp := []byte{}
		for j := 0; j < 2*(n-i); j++ {
			tmp = append(tmp, ' ')
		}
		for j := 0; j <= i; j++ {
			tmp = append(tmp, byte('0'+j))
			tmp = append(tmp, ' ')
		}
		for j := i - 1; j >= 0; j-- {
			tmp = append(tmp, byte('0'+j))
			tmp = append(tmp, ' ')
		}
		s = append(s, string(tmp[:len(tmp)-1]))
	}
	for _, t := range s {
		fmt.Println(t)
	}
	for i := len(s) - 2; i >= 0; i-- {
		fmt.Println(s[i])
	}

}
