package serach

import (
	"fmt"
)

func linearSerach(arr []int, v int) bool {
	fmt.Println("Searcing:", arr, "For:", v)
	found := false
	for i := 0; i < len(arr); i++ {
		if arr[i] == v {
			found = true
		}
	}

	return found
}