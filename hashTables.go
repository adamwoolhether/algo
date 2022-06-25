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

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Exercises

// intersectionOfTwoArrays will return an array of all strings that are within both of the given arrays.
func intersectionOfTwoArrays(arr1, arr2 []string) []string {
	var smallerArray []string
	var largerArray []string
	var intersections []string

	if len(arr1) > len(arr2) { // no check if same size...
		smallerArray = arr2
		largerArray = arr1
	} else {
		smallerArray = arr1
		largerArray = arr2
	}

	hashTable := make(map[string]struct{}, len(smallerArray))

	for _, v := range smallerArray {
		hashTable[v] = struct{}{}
	}

	for _, v := range largerArray {
		if _, ok := hashTable[v]; ok {
			intersections = append(intersections, v)
		}
	}

	return intersections
}

// firstDuplicate returns the first duplicated string in the slice, or a blank string if no duplicates are found.
func firstDuplicate(arr []string) string {
	hashMap := make(map[string]struct{})

	for _, v := range arr {
		if _, ok := hashMap[v]; ok {
			return v
		}
		hashMap[v] = struct{}{}
	}

	return ""
}

// missingLetter will return the first letter of the alphabet that isn't used in the given string.
func missingLetter(str string) string {
	hashMap := make(map[uint8]struct{})
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	for i := 0; i < len(str); i++ {
		hashMap[str[i]] = struct{}{}
	}

	for i := 0; i < len(alphabet); i++ {
		if _, ok := hashMap[alphabet[i]]; !ok {
			return string(alphabet[i])
		}
	}

	return ""
}

// firstNonDuplicated will return the first NON duplicated char in a given string.
func firstNonDuplicated(str string) string {
	hashMap := make(map[uint8]int)

	for i := 0; i < len(str); i++ {
		if _, ok := hashMap[str[i]]; !ok {
			hashMap[str[i]] = 1
		} else {
			hashMap[str[i]]++
		}
	}

	for i := 0; i < len(str); i++ {
		if count := hashMap[str[i]]; count == 1 {
			return string(str[i])
		}
	}

	return ""
}
