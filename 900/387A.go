package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s, t string
	fmt.Fscan(in, &s, &t)

	tt := strings.Split(t, ":")
	ss := strings.Split(s, ":")
	h, _ := strconv.Atoi(tt[0])
	m, _ := strconv.Atoi(tt[1])

	hh, _ := strconv.Atoi(ss[0])
	mm, _ := strconv.Atoi(ss[1])

	if mm < m {
		mm += 60
		h++
	}
	mm -= m
	if hh < h {
		hh += 24
	}
	hh -= h

	fmt.Fprintf(out, "%02d:%02d\n", hh, mm)
}
