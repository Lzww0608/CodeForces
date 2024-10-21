package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	s := [2]string{}
	t := [2]string{}
	for i := 0; i < 2; i++ {
		fmt.Fscan(in, &s[i])
	}
	for i := 0; i < 2; i++ {
		fmt.Fscan(in, &t[i])
	}
	a := s[0] + string(s[1][1]) + string(s[1][0])
	seq := []byte{}
	for i := 0; len(seq) < 3; i++ {
		if a[i%4] == 'A' || (len(seq) > 0 && a[i%4] != 'X') {
			seq = append(seq, a[i%4])
		}
	}
	a = t[0] + string(t[1][1]) + string(t[1][0])
	seq1 := []byte{}
	for i := 0; len(seq1) < 3; i++ {
		if a[i%4] == 'A' || (len(seq1) > 0 && a[i%4] != 'X') {
			seq1 = append(seq1, a[i%4])
		}
	}
	for i := 0; i < 3; i++ {
		if seq[i] != seq1[i] {
			fmt.Fprintln(out, "NO")
			return
		}
	}
	fmt.Fprintln(out, "YES")
}
