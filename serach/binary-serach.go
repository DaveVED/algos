package main

import (
	"fmt"
)

func binarySerach(arr []int, needle int) bool {	
	found := false
	low := 0
	high := len(arr)
	
	for low < high {
		mid := low + (high - low)/2
		value := arr[mid]

		if value == needle {
			found = true
			break;
		} else if value > needle {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return found
}

func main() {
	arr := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	v := 6
	found := binarySerach(arr, v)
	fmt.Println(found)

	v = 16
	found = binarySerach(arr, v)
	fmt.Println(found)
}