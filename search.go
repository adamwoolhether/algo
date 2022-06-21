package algo

// linearSearch searches a given slice for the presence of searchValue,
// returning the index of the found searchValue if found.
// O(N)
func linearSearch[T numbers](arr []T, searchValue T) int {
	for i, value := range arr {
		if value == searchValue {
			return i
		}

		if value > searchValue {
			break
		}
	}

	return -1
}

// binarySearch searches a given slice for the presence of searchValue,
// returning the index of the found searchValue if found.
// O(log N)
func binarySearch[T numbers](arr []T, searchValue T) int {
	lower := 0
	upper := len(arr) - 1

	for lower <= upper {
		mid := (lower + upper) / 2
		midVal := arr[mid]

		switch {
		case searchValue == midVal:
			return mid
		case searchValue < midVal:
			upper = mid - 1
		case searchValue > midVal:
			lower = mid + 1
		}

	}

	return -1
}
