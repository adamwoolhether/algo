package algo

// isSubsetInefficient determines if a smaller array is a subset of another larger array.
// This implementation doesn't use a hash table.
// O(N * M)
func isSubsetInefficient(arr1 []string, arr2 []string) bool {
	var smallerArray []string
	var largerArray []string

	if len(arr1) > len(arr2) { // no check if same size...
		smallerArray = arr2
		largerArray = arr1
	} else {
		smallerArray = arr1
		largerArray = arr2
	}

	for _, smallArrNum := range smallerArray {
		match := false
		for _, largeArrNum := range largerArray {
			if smallArrNum == largeArrNum {
				match = true
				break
			}
		}

		if !match {
			return false
		}
	}

	return true
}

// isSubset determines if a smaller array is a subset of another larger array.
// It uses a hashtable for better performance.
// O(N)
func isSubset(arr1 []string, arr2 []string) bool {
	var smallerArray []string
	var largerArray []string

	if len(arr1) > len(arr2) { // no check if same size...
		smallerArray = arr2
		largerArray = arr1
	} else {
		smallerArray = arr1
		largerArray = arr2
	}

	hashTable := make(map[string]struct{}, len(largerArray))

	for _, v := range largerArray {
		hashTable[v] = struct{}{}
	}

	for _, v := range smallerArray {
		if _, ok := hashTable[v]; !ok {
			return false
		}
	}

	return true
}
