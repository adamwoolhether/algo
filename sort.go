package algo

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
