package algo

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/exp/constraints"
)

// twoSum1 determines if the sum of any two numbers in a given list
// will add up to the target number.
// It demonstrates an inefficient solution with an O(N^2) runtime.
func twoSumInefficient[T numbers](nums []T, target T) bool {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ { // compare counterparts
			if i != j && nums[i]+nums[j] == target {
				return true
			}
		}
	}

	return false
}

// twoSum determines if the sum of any two numbers in a given list
// will add up to the target number.
// It optimizes the twoSumInefficient solution with an O(N) runtime.
func twoSum[T numbers](nums []T, target T) bool {
	numMap := make(map[T]struct{}, len(nums))

	for _, num := range nums {
		// Check if counterpart to num is in the map.
		if _, ok := numMap[target-num]; ok {
			return true
		}
		// Store each number as a key in the map.
		numMap[num] = struct{}{}
	}

	return false
}

// gameWinnerInefficient determines who will be the winner of The Coin Game.
// It's an inefficient implementation with O(2^n) speed.
func gameWinnerInefficient(numberOfCoins int, currentPlayer string) string {
	var nextPlayer string

	if numberOfCoins <= 0 {
		return currentPlayer
	}

	if currentPlayer == "you" {
		nextPlayer = "them"
	} else if currentPlayer == "them" {
		nextPlayer = "you"
	}

	if gameWinnerInefficient(numberOfCoins-1, nextPlayer) == currentPlayer ||
		gameWinnerInefficient(numberOfCoins-2, nextPlayer) == currentPlayer {
		return currentPlayer
	} else {
		return nextPlayer
	}
}

// gameWinner is an efficient algorith to determine the winner of The Coin Game.
// O(1)
func gameWinner(numberOfCoins int) string {
	if (numberOfCoins-1)%3 == 0 {
		return "them"
	}

	return "you"
}

// determines the indices of two slices that, if switched
// would give the two arrays the same sum. An empty array is
// returned if non are found.
// It's an inefficient algorithm, using nested loops.
func sumSwapInefficient[T numbers](slice1 []T, slice2 []T) [2]int {
	sum := func(nums []T) T {
		var sum T
		for _, v := range nums {
			sum += v
		}

		return sum
	}

	sum1 := sum(slice1)
	sum2 := sum(slice2)

	var newSum1 T
	var newSum2 T
	for i := 0; i < len(slice1); i++ {
		for j := 0; j < len(slice2); j++ {
			newSum1 = sum1 - slice1[i] + slice2[j]
			newSum2 = sum2 - slice2[j] + slice1[i]
			if newSum1 == newSum2 {
				return [2]int{i, j}
			}
		}
	}

	return [2]int{}
}

// sumSwap determines the indices of two slices that, if switched
// would give the two arrays the same sum. An empty array is
// returned if non are found.
// Time Complexity: O(N + M)
func sumSwap[T numbers](slice1 []T, slice2 []T) [2]int {
	hashMap := make(map[T]int)
	var sum1 T
	var sum2 T

	// Get sum of first array, storing its values in a hash table with its index.
	for i, v := range slice1 {
		sum1 += v
		hashMap[v] = i
	}

	// Get sum of second array
	for _, v := range slice2 {
		sum2 += v
	}

	// Calculate how much the array sums must shift to be equal.
	shiftAmount := (sum1 - sum2) / 2

	// Iterate over each number in the second array
	for index2, v := range slice2 {
		// Check hashmap for the number's counterpart.
		if index1, ok := hashMap[v+shiftAmount]; ok {
			return [2]int{index1, index2}
		}
	}

	return [2]int{}
}

// maxGreedy finds the greatest number in a slice.
// It represents a greedy version of the max algorithm
func maxGreedy[T numbers](nums []T) T {
	greatest := nums[0]

	for _, num := range nums {
		if num > greatest {
			greatest = num
		}
	}

	return greatest
}

// maxSum returns the greatest sum of any contiguous subsection
// in a given slice of numbers.
// Time Complexity: O(1)
func maxSum[T numbers](nums []T) T {
	var currentSum T
	var greatestSum T

	for i, number := range nums {
		fmt.Println(i, number, currentSum, greatestSum)
		// If current number is < 0, reset the current sum to zero.
		if currentSum+number < 0 {
			currentSum = 0
			continue
		}
		currentSum += number

		// Greedily assume current sum is the greatest sum
		// if it's the greatest sum we've encountered so far.
		if currentSum > greatestSum {
			greatestSum = currentSum
		}
	}

	return greatestSum
}

// increasingTriplets will determin if there is an upward trend in a given slice of
// prices. An 'uptrend' is defined as three increasing prices.
func increasingTriplets(stockPrices []float64) bool {
	lowestPrice := stockPrices[0]
	middlePrice := math.Inf(1)

	for _, price := range stockPrices {
		if price <= lowestPrice {
			lowestPrice = price
		} else if price <= middlePrice { // If current price > lowest, but lower than middle.
			middlePrice = price
		} else { // Current price is higher than middle price.
			return true
		}
	}

	return false
}

// areAnagramsNested is a slightly better approach than the anagramsOf implementation.
// It runs at O(N*M). ***Will need to analyze Time Complexity of ReplaceAll to be more accurate.
func areAnagramsNested(firstString, secondString string) bool {
	firstString = strings.ToLower(strings.ReplaceAll(firstString, " ", ""))
	secondString = strings.ToLower(strings.ReplaceAll(secondString, " ", ""))

	// Convert second string into an array, allowing us to delete chars from it.
	for i := 0; i < len(firstString); i++ {
		// If we're iterating through the first string but the second string is empty.
		if len(secondString) == 0 {
			return false
		}

		for j := 0; j < len(secondString); j++ {
			// If the same char found in both strings.
			if firstString[i] == secondString[j] {
				// Delete the char from the second array and return to outer loop.
				secondString = fmt.Sprintf("%s%s", secondString[:j], secondString[j+1:])
				break
			}

			if j == len(secondString)-1 { // Indicates that secondString lacks a char that firstString has.
				return false

			}
		}
	}

	// If there are no chars remaining in the secondString
	// after iterating over the first, then the two strings are an anagram.
	return len(secondString) == 0
}

// areAnagramsSorted first sorts the two strings, comparing them
// side-by-side to determine if they're an anagram of each other.
// Theoretical time: O(N log N + M log M)
func areAnagramsSorted(firstString, secondString string) bool {
	firstRunes := []rune(strings.ToLower(strings.ReplaceAll(firstString, " ", "")))
	secondRunes := []rune(strings.ToLower(strings.ReplaceAll(secondString, " ", "")))
	sort.Slice(firstRunes, func(i int, j int) bool {
		return firstRunes[i] < firstRunes[j]
	})
	sort.Slice(secondRunes, func(i int, j int) bool {
		return secondRunes[i] < secondRunes[j]
	})

	if len(firstRunes) != len(secondRunes) {
		return false
	}

	for i := 0; i < len(firstRunes); i++ {
		if firstRunes[i] != secondRunes[i] {
			return false
		}
	}

	return true
}

// areAnagrams determines if two strings are anagrams by places each
// char in a hash map and comparing the resulting hash maps' number of each char.
// It takes N+M steps.
func areAnagrams(firstString, secondString string) bool {
	firstString = strings.ToLower(strings.ReplaceAll(firstString, " ", ""))
	secondString = strings.ToLower(strings.ReplaceAll(secondString, " ", ""))
	firstStringMap := make(map[uint8]int)
	secondStringMap := make(map[uint8]int)

	// Create a hashmap out of the first string.
	for i := 0; i < len(firstString); i++ {
		if firstString[i] == 32 {
			continue
		}
		firstStringMap[firstString[i]]++
	}
	// Create a hashmap out of the second string.
	for i := 0; i < len(secondString); i++ {
		if secondString[i] == 32 {
			continue
		}
		secondStringMap[secondString[i]]++
	}

	return cmp.Equal(firstStringMap, secondStringMap)
}

// groupArray takes a slice of a type that can be ordered and groups similar values.
// O(N) time.
func groupArray[T constraints.Ordered](array []T) []T {
	hashMap := make(map[T]int)
	result := make([]T, 0, len(array))

	// Store values of each string in a hash table.
	for _, v := range array {
		hashMap[v]++
	}

	// Iterate over the hashmap and populate the result slice
	// with the correct number of each value.
	for val, count := range hashMap {
		for i := 0; i < count; i++ {
			result = append(result, val)
		}
	}

	return result
}

// groupArraySort  takes a slice of a type that can be ordered and groups similar values.
// It uses Go's sort.Slice Method, and is O(log N)
func groupArraySort[T constraints.Ordered](array []T) []T {
	sort.Slice(array, func(i, j int) bool {
		return array[i] < array[j]
	})

	return array
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Exercises

// 1

type Athletes struct {
	FirstName string
	LastName  string
	Team      string
}

// multisportAthletes will take two lists of athletes for a certain sport, returning
// a list of athletes that play on both sports.
// O(N+M)
func findMultisportAthletes(array1, array2 []Athletes) []string {
	athletesMap := make(map[string]struct{}, len(array1))
	multisportAthletes := []string(nil)

	for _, athlete := range array1 {
		fullName := fmt.Sprintf("%s %s", athlete.FirstName, athlete.LastName)
		athletesMap[fullName] = struct{}{}
	}

	for _, athlete := range array2 {
		fullName := fmt.Sprintf("%s %s", athlete.FirstName, athlete.LastName)
		if _, ok := athletesMap[fullName]; ok {
			multisportAthletes = append(multisportAthletes, fullName)
		}
	}

	return multisportAthletes
}

// findMultisportAthletesInefficient will take two lists of athletes for a certain sport, returning
// a list of athletes that play on both sports.
// O(N * M)
func findMultisportAthletesInefficient(array1, array2 []Athletes) []string {
	multisportAthletes := []string(nil)

	for _, athlete1 := range array1 {
		name1 := fmt.Sprintf("%s %s", athlete1.FirstName, athlete1.LastName)
		for _, athlete2 := range array2 {
			name2 := fmt.Sprintf("%s %s", athlete2.FirstName, athlete2.LastName)
			if name1 == name2 {
				multisportAthletes = append(multisportAthletes, name1)
			}
		}
	}

	return multisportAthletes
}

// 2

// findMissingNum will determine if any number in an incremented
// array of numbers is missing. compare this with findMissingNumber
// in partitioning.go, which uses sort.
// O(N)
func findMissingNum(array []int) int {
	shouldBe, actualSum := 0, 0
	for i, v := range array {
		shouldBe += i + 1
		actualSum += v
	}

	diff := shouldBe - actualSum

	return diff
}

// 3
// findGreatestPrice will determine the greatest profit possible to make a single
// buy and sell transaction for a stock.
// O(N)
func findGreatestProfit(prices []int) int {
	buyPrice := prices[0]
	greatestProfit := 0

	for _, price := range prices {
		potentialProfit := price - buyPrice

		if price < buyPrice {
			buyPrice = price
		} else if potentialProfit > greatestProfit {
			greatestProfit = potentialProfit
		}
	}

	return greatestProfit
}

// greatestProduct will find the greatest possible product of any two numbers in an array.
// O(N)
func greatestProduct(nums []int) int {
	greatest := -math.MaxInt64
	secondGreatest := -math.MaxInt64

	lowest := math.MaxInt64
	secondLowest := math.MaxInt64

	for _, num := range nums {
		if num >= greatest {
			secondGreatest = greatest
			greatest = num
		} else if num > secondGreatest {
			secondGreatest = num
		}

		if num <= lowest {
			secondLowest = lowest
			lowest = num
		} else if num > lowest && num < secondLowest {
			secondLowest = num
		}
	}

	greatestProductHighest := greatest * secondGreatest
	greatestProductLowest := lowest * secondLowest

	if greatestProductHighest > greatestProductLowest {
		return greatestProductHighest
	}

	return greatestProductLowest
}

// sortTemperatures will sort a list of given temperatures in O(N) time.
// We can do this because we know the range of temps is 97.0-99.0
func sortTemperatures(temps []float64) []float64 {
	hashMap := make(map[float64]int)

	// Populate the map with occurrences of temps.
	for _, temp := range temps {
		hashMap[temp]++
	}

	sortedArray := make([]float64, 0, len(temps))

	// Multiply temps by 10 in order to increment temp by a whole number during
	// the loop to avoid floating-point errors.
	// NOTE: is this needed?
	temperature := float64(970)

	for temperature <= 990 {
		floatingTemp := temperature / 10.0
		// Check if the map contains the current temp.
		if v, ok := hashMap[floatingTemp]; ok {
			// Populate the sortedArray with as many occurences of the current table.
			for i := 0; i < v; i++ {
				sortedArray = append(sortedArray, floatingTemp)
			}
		}
		temperature++
	}

	return sortedArray
}

// longestSequenceLength finds the length of the longest sequence of numbers
// in a given array.
// O (3N) = O(N)
func longestSequenceLength(nums []int) int {
	hashMap := make(map[int]bool, len(nums))
	greatestSequenceLength := 0

	// Populate hash map with nubers as keys.
	for _, num := range nums {
		hashMap[num] = true
	}

	// Iterate over each number in the array.
	for _, num := range nums {
		// Determine if the number is the first in the sequence(there isn't a lower number).
		if _, ok := hashMap[num-1]; !ok {
			// Start counting the length of the current sequence,
			// starting with current number. As the first number,
			// the length starts at 1.
			currentLength := 1

			// Establish a current number to use in the loop below.
			currentNumber := num

			// Run loop as long as there's a next number in the sequence.
			for hashMap[currentNumber+1] {
				// Move on to next num in sequence.
				currentNumber++

				// Increase length of sequence by 1.
				currentLength++

				// Greedily keep track of greatest sequence length.
				if currentLength > greatestSequenceLength {
					greatestSequenceLength = currentLength
				}
			}
		}
	}

	return greatestSequenceLength
}
