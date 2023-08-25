package main

import (
	"fmt"
	"math"
)

// given two crystal balls that will break if dropped from high enough
// distance, determine the exact spot in which it will break in the most
// optimized way.

func balls(breaks []bool) int {
	jumpAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))
	i := jumpAmount
	for ; i < len(breaks); i += jumpAmount {
		if breaks[i] {
			break
		}
	}
	i = i - jumpAmount
	for j := 0; j < jumpAmount && i < len(breaks); i += 1 {
		if breaks[i] {
			return i
		}
		j++
	}
	return -1
}

func main() {
	arr := []bool{false, false, false, false, false, true, true}
	fmt.Println(balls(arr))
}