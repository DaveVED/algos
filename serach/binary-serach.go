package serach

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