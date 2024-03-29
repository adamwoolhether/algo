package algo

// bubbleSort O(N^2)
func bubbleSort[T numbers](arr []T) []T {
	sorted := false

	for !sorted {
		sorted = true

		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
		}
	}

	return arr
}

// selectionSort O(N^2)
func selectionSort[T numbers](arr []T) []T {
	for i := 0; i < len(arr)-1; i++ {
		lowestNumIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[lowestNumIdx] {
				lowestNumIdx = j
			}
		}

		if lowestNumIdx != i {
			temp := arr[i]
			arr[i] = arr[lowestNumIdx]
			arr[lowestNumIdx] = temp
		}
	}

	return arr
}

func insertionSort[T numbers](arr []T) []T {
	for i := 1; i < len(arr); i++ {
		tempVal := arr[i]
		position := i - 1

		for position >= 0 {
			if arr[position] > tempVal {
				arr[position+1] = arr[position]
				position--
			} else {
				break
			}
		}

		arr[position+1] = tempVal
	}

	return arr
}
