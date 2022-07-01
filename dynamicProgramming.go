package algo

// max finds the greatest number from an array {
func max(slice []int) int {
	if len(slice) == 1 {
		return slice[0]
	}

	maxRemainder := max(slice[1:])

	if slice[0] > maxRemainder {
		return slice[0]
	}

	return maxRemainder
}

func maxInefficient(slice []int) int {
	if len(slice) == 1 {
		return slice[0]
	}

	if slice[0] > maxInefficient(slice[1:]) {
		return slice[0]
	}

	return maxInefficient(slice[1:])
}

func fibInefficient(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return fibInefficient(n-2) + fibInefficient(n-1)
}

// fibMemoization implements a recursive fibonacci function with memoization
// Because this is O(2N) - 1, we can set the size of the map we give:
// memo := make(map[int]int, 2*n-1)
func fibMemoization(num int, memo map[int]int) int {
	if num == 0 || num == 1 {
		return num
	}

	if _, recorded := memo[num]; !recorded {
		memo[num] = fibMemoization(num-2, memo) + fibMemoization(num-1, memo)
	}

	return memo[num]
}

// fib also uses memoization and a closure to run recursion.
func fib(number int) int {
	memo := make(map[int]int, 2*number-1)

	var recurse func(n int, mem map[int]int) int
	recurse = func(n int, mem map[int]int) int {
		if n == 0 || n == 1 {
			return n
		}

		if _, recorded := mem[n]; !recorded {
			mem[n] = recurse(n-2, mem) + recurse(n-1, mem)
		}

		return mem[n]
	}

	return recurse(number, memo)
}

func fibBottomUp(n int) int {
	if n <= 1 {
		return n
	}

	a, b := 0, 1

	for i := 2; i <= n; i++ {
		next := a + b
		a = b
		b = next
	}

	return b
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Exercises
// addUntil100 returns the sum of an []int, ignoring the number
// if adding it to the current sum exceeds 100.
func addUntil100(ints []int) int {
	if len(ints) == 0 {
		return 0
	}

	sumOfRemainders := addUntil100(ints[1:])

	if ints[0]+sumOfRemainders > 100 {
		return sumOfRemainders
	}

	return ints[0] + sumOfRemainders
}

// golomb calculates the nth number from a "Golomb Sequence"
// https://en.wikipedia.org/wiki/Golomb_sequence
func golomb(number int) int {
	memo := make(map[int]int)

	var recurse func(n int, mem map[int]int) int
	recurse = func(n int, mem map[int]int) int {
		if n == 1 {
			return 1
		}

		if _, recorded := mem[n]; !recorded {
			mem[n] = 1 + recurse(n-recurse(recurse(n-1, mem), mem), mem)
		}

		return mem[n]
	}
	return recurse(number, memo)
}

func uniquePathsMemo(rows, columns int) int {
	memo := make(map[[2]int]int)

	var recurse func(r, c int, mem map[[2]int]int) int
	recurse = func(r, c int, mem map[[2]int]int) int {
		if r == 1 || c == 1 {
			return 1
		}

		if _, recorded := mem[[2]int{r, c}]; !recorded {
			mem[[2]int{r, c}] = recurse(r-1, c, mem) + recurse(r, c-1, mem)
		}

		return mem[[2]int{r, c}]
	}

	return recurse(rows, columns, memo)
}
