package algo

import "fmt"

func averageOfEvenNumbers(arr []int) int {
	sum := 0
	count := 0

	for _, num := range arr {
		if num%2 == 0 {
			sum += num
			count++
		}
	}

	return sum / count
}

func wordBuilder(arr []string) []string {
	var collection []string

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i != j {
				collection = append(collection, arr[i]+arr[j])
			}
		}
	}

	return collection
}

func sample(arr []int) (int, int, int) {
	first := arr[0]
	middle := arr[len(arr)/2]
	last := arr[len(arr)-1]

	return first, middle, last
}

func markInventory(clothingItems []string) []string {
	clothingOptions := []string{}

	for _, item := range clothingItems {
		for i := 1; i < 6; i++ {
			clothingOptions = append(clothingOptions, fmt.Sprintf("%s Size: %d", item, i))
		}
	}

	return clothingOptions
}

func countOnes(outerArr [][]int) int {
	count := 0

	for _, innerArray := range outerArr {
		for _, num := range innerArray {
			if num == 1 {
				count++
			}
		}
	}

	return count
}

func isPalindrome(str string) bool {
	leftIndex := 0
	rightIndex := len(str) - 1

	// Interate until left index reaches middle of the array.
	for leftIndex < len(str)/2 {
		if str[leftIndex] != str[rightIndex] {
			return false
		}
		leftIndex++
		rightIndex--
	}

	return true
}

// twoNumbers returns the product of every combination of two numbers in an array.
func twoNumbers(arr []int) []int {
	var products []int

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			products = append(products, arr[i]*arr[j])
		}
	}

	return products
}

// twoNumberProducts computes the product of every number from one array by every number of a second array.
func twoNumberProducts(arr1 []int, arr2 []int) []int {
	var products []int

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			products = append(products, arr1[i]*arr2[j])
		}
	}

	return products
}

// ruby code:
// func everyPassword(n int) {
// 	for l := 'a'; l < 'z'; l++ {
// 		for t := 0; t < n; t++ {
// 			fmt.Print(l)
// 		}
// 	}
// }

// oneHundredSum returns true if a given array is a '100-sum-arry'
// A '100-sum-arry' has the following criteria:
// - First and last numbers add up to 100
// - Second and second-to-last numbers adds up to 100
// - Third and third-to-last numbers add up to 100, and so on.
// O(N/2) or O(N)
func oneHundredSum(arr []int) bool {
	leftIndex := 0
	rightIndex := len(arr) - 1

	for leftIndex < len(arr)/2 {
		if arr[leftIndex]+arr[rightIndex] != 100 {
			return false
		}
		leftIndex++
		rightIndex--
	}

	return true
}

// merge takes two arrays and combines them.
// O(N + M), or O(N)
func merge(arr1 []int, arr2 []int) []int {
	result := make([]int, len(arr1)+len(arr2))
	arr1Pointer := 0
	arr2Pointer := 0

	// Run until the end of both arrays has been reached.
	for arr1Pointer < len(arr1) || arr2Pointer < len(arr2) {
		switch {
		case arr1[arr1Pointer] >= len(arr1):
			// If end of first array is reached, add item from second array.
			result = append(result, arr2[arr2Pointer])
			arr2Pointer++
		case arr2[arr2Pointer] >= len(arr2):
			// If end of second array reached, add item from first array.
			result = append(result, arr1[arr1Pointer])
			arr1Pointer++
		case arr1[arr1Pointer] < arr2[arr2Pointer]:
			// If the current number in first array is less than the current num in second array.
			result = append(result, arr1[arr1Pointer])
			arr1Pointer++
		default:
			// If current number in second array is less than the current num in first array.
			result = append(result, arr2[arr2Pointer])
			arr2Pointer++
		}
	}

	return result
}

// findNeedle determines if a string, 'needle', is a substring of 'haystack'.
// O(N * M)
func findNeedle(needle, haystack string) bool {
	needleIdx := 0
	haystackIdx := 0

	for haystackIdx < len(haystack) {
		if needle[needleIdx] == haystack[haystackIdx] {
			foundNeedle := true

			for needleIdx < len(needle) {
				if needle[needleIdx] != haystack[haystackIdx+needleIdx] {
					foundNeedle = false
					break
				}
				needleIdx++
			}
			if foundNeedle {
				return true
			}
			needleIdx = 0
		}
		haystackIdx++
	}

	return false
}

// largestProduct finds the greatest product of three numbers in a given array.
// O(N^3)
func largestProduct(arr []int) int {
	largestProductSoFar := arr[0] * arr[1] * arr[2]
	i := 0

	for i < len(arr) {
		j := i + 1

		for j < len(arr) {
			k := j + 1

			for k < len(arr) {
				if product := arr[i] * arr[j] * arr[k]; product > largestProductSoFar {
					largestProductSoFar = product
				}
				k++
			}
			j++
		}
		i++
	}

	return largestProductSoFar
}

// pickResumes cuts an array of resumes in half until there is only one left,
// alternating between removing the bottom or the top half.
// O(log N)
func pickResumes(resumes []string) string {
	eliminate := "top"

	for len(resumes) > 1 {
		// fmt.Printf("Len: %d\tCap: %d\tSlice: %v\n", len(resumes), cap(resumes), resumes)
		if eliminate == "top" {
			resumes = append(resumes[:len(resumes)/2], resumes[len(resumes)/2:len(resumes)/2]...)
			eliminate = "bottom"
		} else {
			resumes = append(resumes[len(resumes)/2:], resumes[len(resumes):len(resumes)]...)
			eliminate = "top"
		}
	}

	return resumes[0]
}
