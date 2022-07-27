package algo

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
