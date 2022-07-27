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
