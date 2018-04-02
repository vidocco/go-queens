package main

import (
	"flag"
	"math"
	"time"
)

func main() {
	size := flag.Int("size", 0, "board size for the n-queens problem")

	flag.Parse()
	println("boardsize: ", *size)
	start := time.Now()
	println("solutions found: ", queens(*size, 0, 0, 0, *size-1, *size-1))
	println("time elapsed (us): ", time.Since(start))
}

func queens(size int, row int, R int, L int, pos int, level int) int {
	count, answer, limit := 0, 0, math.Pow(2, float64(size))
	R, L = R>>1, L<<1
	if float64(L) >= limit {
		L -= int(limit)
	}
	valid := R | L | row
	for i := 0; i < size; i++ {
		if (valid & int(math.Pow(2, float64(i)))) == 0 {
			count++
		}
	}
	if (level == 0 && count == 1) || (level != 0 && count == 0) {
		return (count)
	}
	for j := pos; j >= 0; j-- {
		curr := int(math.Pow(2, float64(j)))
		if curr&valid == 0 {
			prevRow, prevL, prevR := row, L, R
			row, L, R = curr|prevRow, L|curr, R|curr
			answer += queens(size, row, R, L, pos, (level - 1))
			row, L, R = prevRow, prevL, prevR
		}
	}
	return answer
}
