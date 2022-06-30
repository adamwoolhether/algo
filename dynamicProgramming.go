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
func fibMemoization(n int, memo map[int]int) int {
	if n == 0 || n == 1 {
		return n
	}

	if _, recorded := memo[n]; !recorded {
		memo[n] = fibMemoization(n-2, memo) + fibMemoization(n-1, memo)
	}

	return memo[n]
}

// fib also uses memoization and a closure to run recursion.
func fib(n int) int {
	memo := make(map[int]int, 2*n-1)

	var recurse func(num int, mem map[int]int) int
	recurse = func(num int, mem map[int]int) int {
		if num == 0 || num == 1 {
			return num
		}

		if _, recorded := mem[num]; !recorded {
			mem[num] = recurse(num-2, mem) + recurse(num-1, mem)
		}

		return mem[num]
	}

	return recurse(n, memo)
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
func golomb(n int) int {
	memo := make(map[int]int)

	var recurse func(num int, m map[int]int) int
	recurse = func(num int, m map[int]int) int {
		if num == 1 {
			return 1
		}

		if _, recorded := m[num]; !recorded {
			m[num] = 1 + recurse(num-recurse(recurse(num-1, m), m), m)
		}

		return m[num]
	}
	return recurse(n, memo)
}

func uniquePathsMemo(rows, columns int) int {
	memo := make(map[[2]int]int)

	var recurse func(r, c int, m map[[2]int]int) int
	recurse = func(r, c int, m map[[2]int]int) int {
		if r == 1 || c == 1 {
			return 1
		}

		if _, recorded := m[[2]int{r, c}]; !recorded {
			m[[2]int{r, c}] = recurse(r-1, c, m) + recurse(r, c-1, m)
		}

		return m[[2]int{r, c}]
	}

	return recurse(rows, columns, memo)
}
