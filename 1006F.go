package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var H, W int
	var K int64
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	H, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	W, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	K, _ = strconv.ParseInt(scanner.Text(), 10, 64)

	A := make([][]int64, H)
	for y := range A {
		A[y] = make([]int64, W)
		for x := range A[y] {
			scanner.Scan()
			A[y][x], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		}
	}

	V := make([][][][]int64, H)
	for y := range V {
		V[y] = make([][][]int64, W)
		for x := range V[y] {
			V[y][x] = make([][]int64, 2)
			for i := range V[y][x] {
				V[y][x][i] = make([]int64, 0)
			}
		}
	}

	V[0][0][0] = append(V[0][0][0], A[0][0])
	V[H-1][W-1][1] = append(V[H-1][W-1][1], A[H-1][W-1])

	step := H + W - 2
	L := step / 2
	R := step - L

	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			if y+x < L {
				if y+1 < H {
					for _, v := range V[y][x][0] {
						V[y+1][x][0] = append(V[y+1][x][0], v^A[y+1][x])
					}
				}
				if x+1 < W {
					for _, v := range V[y][x][0] {
						V[y][x+1][0] = append(V[y][x+1][0], v^A[y][x+1])
					}
				}
			}
		}
	}

	for y := H - 1; y >= 0; y-- {
		for x := W - 1; x >= 0; x-- {
			if (H-1-y)+(W-1-x) < R {
				if y > 0 {
					for _, v := range V[y][x][1] {
						V[y-1][x][1] = append(V[y-1][x][1], v^A[y-1][x])
					}
				}
				if x > 0 {
					for _, v := range V[y][x][1] {
						V[y][x-1][1] = append(V[y][x-1][1], v^A[y][x-1])
					}
				}
			}
		}
	}

	var ret int64
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			if len(V[y][x][0]) > 0 && len(V[y][x][1]) > 0 {
				M := make(map[int64]int64)
				for _, v := range V[y][x][0] {
					M[v]++
				}
				for _, v := range V[y][x][1] {
					ret += M[K^v^A[y][x]]
				}
			}
		}
	}

	fmt.Println(ret)
}
