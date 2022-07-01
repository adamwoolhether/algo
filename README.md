**Algorithm**: a set of instructions for completing a specific task.

# Big O
#### The Key Question
Big O helps us answer the **Key Question:**  
🔑 <span style="font-size:larger;">_If there are N data elements, how many steps will the algorithm take?_</span>

#### Soul of Big O
Big O is much more than just how many steps an algorithm takes. It tells the story of how the number of steps increase.  
🛟 <span style="font-size:larger;">_How will an algorithm's performance change as the data increases?_</span>

Why is this important?  
Suppose we have two algorithms. One O(1) algo takes 100 steps, regardless of input size. The other, a linear O(N). Which is faster?  
_For all arrays greater than 100, the O(N) algo takes longer._

* Big O notation generally refers the the worst-case scenario, unless noted otherwise.


### Speeds
| Big O            |   Pronounciation   | Time Complexity | Notes                   |
|:-----------------|:------------------:|:----------------|:------------------------|
| O(1)             |     "O of one"     | Constant time   |                         |
| O(log N)         |    "O of log N"    | Log time        |                         |
| O(N)             |      "O of N"      | Linear time     | O(log<sub>2</sub> N)    |
| O(N<sup>2</sup>) |  "O of N squared"  | Quadratic time  | Typical of nested loops |
| O(N!)            | "O of factorial N" | Factorial Time  |                         |

### Logarithms
Logarithms are the inverse of _exponents_.  
Exponent example: 2<sup>3</sup> = 2 * 2 * 2 = 8  
The converse of the above exponent is: log<sub>2</sub>8 = 3  
We had to multiply 2 by itself three times to get 8.
We can also look at it as: 8 / 2 / 2 / 2 = 1
#### O (log N)
Put simply: _O(log N) means the algorithm takes as many steps as it takes to keep halving the data elements until we remain with 1._

### Bubble Sort - A quadratic algo
([see code example](sort.go))  
Two significant steps:  _comparisons_ and _swaps_.  
For N elements, we make:  
`(N - 1) + (N - 2) + (N - 3) ... + 1` comparisons.  

Given 10 elements:
Worst case: need to swap for each comparison; 20 total steps (10 swaps + 10 comparisons).  
This grows at approximately O(N<sup>2</sup>)

#### Another Quadratic Example:
```go
func hasDuplicateValue(arr []int) bool {
    for i := 0; i < len(arr); i++ {
        for j := 0; j < len(arr); j++ {
            if i != j && arr[i] == arr[j] {
                return true
            }
        }       
    }
}
```
Refactoring the above func to be linear (This has storage inefficiencies):
```go
func hasDuplicateValue(arr []int) bool {
    existingNumbers := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
        if existingNumbers[i] == 1 {
            return true
        } else {
            existingNumbers[i] = 1
        }      
    }       
}
```

## Optimization
How to compare two algorithms that seem to have the same efficiency?  
We'll use ([Selection Sort](sort.go)), and compare it to Bubble Sort.  
Selection sort will make one or zero swaps, compared to bubble sort which will make a swap for _each_ and every comparison
in the worst case. Given this, Selection Sort can take about half the amount of steps as Bubble Sort, or seemingly O(N<sup>2</sup> / 2).
Why are they the same Big O speed?

👉 _Big O ignores constants._  
For example:
* N / 2 steps == O(N)
* N<sup>2</sup> + 10 == O(N<sup>2</sup>)
* 2N steps == O(N)
* O(1000N) == O(N)

👉 _Big O only cares about general categories of algorithm speed._ This ties back to the soul of Big O.
So, when two algos fall under the same category, further analysis is needed.

#### Significant steps
We count the number of all steps in algos that fall under the same category.

It's import to consider _all_ scenarios when choosing an algorithm(worst/average/best case). As an example, well compare
Insertion Sort with selection sort.

### Insertion Sort
In the worst case, insertion sort is O(N<sup>2</sup>+N)
* N<sup>2</sup> comparisons and shifts
* N - 1 removals
* N - 1 insertions  

This give N<sup>2</sup> + 2N - 2 Step, which is simplified to O(N<sup>2</sup>+N) because:  
👉 _Big O only takes into account the highest order of N when there are multiple orders added together._  
ie: an algo with N<sup>4</sup> + N<sup>3</sup> + N<sup>2</sup> + N steps is simplified to N<sup>4</sup>.

In the worst case, Selection Sort is faster than Insertion Sort. BUT, we must also consider the _**average case scenario.**_

#### Average Case Scenario
Insertion sort recap:
* Worst-case scenario: compare and shift all data. ()
* Best-case scenario: shift _none_ of data, making just on comparison per pass.
* For average case, we'd probably compare and shift only half the data.

#### Comparison

| Algo           | Best Case         | Average Case      | Worst Case        |
|----------------|-------------------|-------------------|-------------------|
| Selection Sort | N<sup>2</sup> / 2 | N<sup>2</sup> / 2 | N<sup>2</sup> / 2 |
| Insertion Sort | N                 | N<sup>2</sup> / 2 | N<sup>2</sup>     |

Therefore, the better algorithm depends on the context. Both perform similarly in the average case.  
If you can assume that the data will be _mostly_ sorted, use Insertion Sort, but if you can assume it will be mostly sorted
in reverse, then use Selection Sort.  
_If you don't know_, the average case applied, and both are equal.

## Determining Efficiency - Examples

To answer the question of "if there are N data elements, how many steps will the algorithm take?":
1. Determing what the "N" data elements are.
2. Determine how many steps the algorithm takes to process these N values.

### Mean Average of Even Numbers
3N + 3  
O(N)
```go
func averageOfEvenNumbers(arr []int) int {
	sum := 0
	count := 0
	
	for _, num := range arr {
		if num % 2 == 0 {
			sum += num
			count++
		}
	}
	
	return sum / count
}
```
In this example, the loop loops over each of the N elements, so it takes at least N steps.
Inside the loop, it checks if the number is even, and if so, performs two more steps (add to sum, and increment count).  
Big O assumes the worst case, so we say it takes 3N steps, 3 steps for each of the N numbers. Outside the loop, there are
another 3 steps performed. Overall, it the algorithm takes a total of **3N + 3**. After ignoring constants, this is **O(N)**.

### Word Builder
3N<sup>2</sup> + 1   
O(N<sup>2</sup>)
```go
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
```
If there were three nested loops, this would be O(N<sup>3</sup>)

### Array Sample
O(1)
```go
func sample(arr []int) (int, int, int) {
	first := arr[0]
	middle := arr[len(arr)/2]
	last := arr[len(arr)-1]

	return first, middle, last
}
```

### Average Celcius Reading
2N  
O(N)
```go
func averageCelcius(fahrenheitReading []float64) float64 {
	var celciusNumbers []float64

	// Convert each reading to celsius and add to array.
	for _, f := range fahrenheitReading {
		celsiusConversion := (f - 32) / 1.8
		celciusNumbers = append(celciusNumbers, celsiusConversion)
	}

	// Get sum of all Celsius numbers
	var sum float64

	for _, c := range celciusNumbers {
		sum += c
	}

	// Return mean average
	return sum / float64(len(celciusNumbers))
}
```

### Clothing Labels
5N   
O(N)
```go
func markInventory(clothingItems []string) []string {
	clothingOptions := []string{}

	for _, item := range clothingItems {
		for i := 1; i < 6; i++ {
			clothingOptions = append(clothingOptions, fmt.Sprintf("%s Size: %d", item, i))
		}
	}
	
	return clothingOptions
}
```
### Count the Ones
O(N)  
Inner loop only runs for as many numbers as ther are _in total_.
```go
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
```

### Palindrome Checker
N/2 + 3  
O(N)
```go
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
```

### Get All the Products
N<sup>2</sup>/2  
O(N<sup>2</sup>)
```go
func twoNumbers(arr []int) []int {
	var products []int

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			products = append(products, arr[i]*arr[j])
		}
	}
	
	return products
}
```

### Multiple Datasets
O(N * M), where N is the size of one array and M is the size of the second.  
O(N * M) can be thought of as in between O(N) and O(N<sup2></sup>)
```go
func twoNumberProducts(arr1 []int, arr2 []int) []int {
	var products []int
	
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			products = append(products, arr1[i]*arr2[j])
		}
	}
	
	return products
}
```

### Password Cracker
O(26<sup>N</sup>)

below is in ruby
```ruby
def every_password(n)
  (("a" * n)..("z" * n)).each do |str|
    puts str
  end
end
```

## HashMaps
Allow O(1) lookup speed of key-value pairs.  
#### Hashing
Conversion of chars into a numbers via a _hash function_.  Lookups done in two steps: 
1. Hashing the key.
2. Look in cell of hashmap if a value exists for the given key.
#### Collisions
When the two keys, after hashing, have the same number. One approach to resolve this is _separate chaining_:
**Separate chaining**: The value becomes an array. (because of this the _worst case_ of a hash lookup is O(N), if all values get placed in the same key.)

## Efficiency
Hash table efficiency depends on:
* Amount of data stored in the hash table
* How many cells are available in the hash table
* The hash function being used.    

A good hash table _strikes a good balance between avoiding collisions and not consuming large amounts of memeory_  
#️⃣ **For every 7 data elements stored in the hash table, it should have 10 cells**   
**Load Factor:** the ratio of data to cells.  
#️⃣ _Ideal load factor:_`0.7` = `7 elements / 10 cells`

## Stacks and Queues
These are simply _arrays with restrictions_, specifically suitable for handling temporary data, with an emphasis on the _order_ data is handled.

### Stacks
A LIFO data structure.   
Three constraints:
* Data is inserted at the end of the stack.
* Data is deleted from the end of the stack.
* Only the last element can be read.

_Pushing onto the stack_: inserting data.  
_Popping from the stack_: removing data.  
_Peaking_: reading the last element without removal.

Stacks are _abstract data types_: a data structure with a theoretical set of rules revlolving around some other concrete 
data structure.  
Abstract data structures help prevent bugs and give us mental models to work with for a problem.

### Queues
A FIFO data structure.  
Three constraints:
* Data is inserted at the end of the queue.
* Data is delete from the fron of the queue.
* Only the first elemen can be read.

_Enqueue_: inserting data.  
_Dequeue_: removing data.  
_Peaking_: reading the last element without removal.

## Recursion
Recursion is when a function calls itself.
Arbitrary loop vs recursive comparison:
```go
// loop
func countdown(n int) {
	for i := n; i >= 0; i-- {
		fmt.Println(i)
    }
}

// recursion
func countdown(n int) {
    fmt.Println(n)
	
	if n == 0 { // The base case
		return
    }
	countdown(n - 1)
}
```
#### Base Case
👉 _Base case_ is the cae in which the function _will not_ recurse. **All recursive functions** need a base case, or we encounter stack overflow.

### Reading recursive code
1. Identify the base case
2. Walk through the function for the base case
3. Identify the "next-to-last" case, the case just before the base case
4. Repeat this process by identifying the case before the one you just analyzed, walking through the function for that case
```go
func factorial(n int) {
    // identify the base case, the number 1
	if n == 1 { 
		return 1
    }
	
	// the next case:
    return n * factorial(n - 1)
}
```
Walking through the above example, when n = 4:  
`factorial(4) returns 24` == `4 * factorial(3)`  
`factorial(3) returns 6` == `3 * factorial(2)`  
`factorial(2) returns 2` == `2 * factorial(1)`  
`factorial(1) returns 1` This is another way to word the base case for this func.  

Simply put: `( 4 * 3 * 2 * 1) = 24`

### How computers view recursion
Func calls are places on the _call-stack_. The return value is passed up through the call stack, to the calling function.

## Writing Recursive Code
There are a few patterns that make a problem suitable for recursion:

### Repeatedly Execute
When the goal of an algorithm is to repeatedly execute a task. The countdown example above encompasses this.
#### Trick: passing extra parameters
Often, having to pass a separate default parameter is cumbersome, so some languages allow a default param.  
Go doesn't allow passing default parameters, so we use an enclosure.
```go
// doubleArray will double each element of a []int in place.
func doubleArray(arr []int) {
	index := 0
	
	var recurse func(arr []int, index int)
	recurse = func(arr[]int, index int) {
		if index == len(arr) {
			return
        }
		
		arr[index] *= 2
		index++
		recurse(arr, index)
    }
	
	recurse(arr, index)
}
```

### Calculations
When the goal is to make a calculatoin based on a sub-problem of the problem at hand.  
_Subproblem_: the very same problem applied to a smaller input.
ex: factorials. `6! = 6 * 5 * 4 * 3 * 2 * 1`. Here, we know that factorial(6) will be multiplied by whatever factorial(5) is.
So `factorial(6)` is the same as `6 * factorial(5)`.

The factorial example above exemplifies this. In that solution, we compute the result as `n` multiplied by the subproblem `factorial(n - 1)`

#### Two approaches to calculations
#### Bottom up
Requires passing extra params. Bottom up is the same strategy used for making loops **or** recursion; the same computational approach.
#### Top down
Making calculations based on the problem's subproblems. _Needs_ recursion.

### Top-Down Recursion
Allows us to mentally "kick the problem down the road.   
We don't have to understand how the function calling works so solve the problem, as in the `return n * factorial(n - 1` statement. 
#### Top-Down Thought Process
1. Imagine the function you're writing has already been implemented before.
2. Identify the subproblem of the problem.
3. See what happens when you call the funciton on the subpoblem and go from there.

#### Examples
_Array Sum_  
A func that sums all nums in an array. Given an array `[1, 2, 3, 4, 5]`  
- Assume it has alraedy been implemented. 
- Identify the subproblem: We can say the _subproblem_ is `[2, 3, 4, 5]`, all numbers except for the first. 
- Try apply the sum func to our sub problem: `sum([2, 3, 4, 5] = 14`. Adding the first number, `1`, to the result.
  - i.e. `return array[0] + sum(array[1:len(array)-1])`, 
- Lastly, be sure to handle the base case! `if len(array) == 1; { return array[0] }`

_String Reversal_  
Given string `"abcde"`, the subproblem is `"bcde`.  
Now pretend the `revers()` func already exists, so we can call `reverse("bcde")` to get `"edcb"`.  
After that, we'd just throw `"a"` at the end.
```go
func reverse(s string) string {
    if len(s) == 0 {
        return ""
    }

    return reverse(s[1:]) + string(s[0])
}
```

_Counting X_  
Return the number of "x's" in a given string.  
Given string `"axbxcxd"` it should reutnr 3. Subproblem is `"xbxcxd"`  
We call `countX("xbxcxd")` and get 3. We would need to add 1 if the first char was "x", or just return that if not.
```go
func countX(str string) int {
    if len(str) == 0 {
        return 0
    }

    if str[0] == 'x' {
        return 1 + countX(str[1:])
    }

    return countX(str[1:])
}
```

### The Staircase Problem
Given a staircase of N steps, you have the ability to climb 1, 2, or 3 steps at a time. 
How many different possible paths can someone take to reach the top?

Ex:  
Two steps - Two possible paths: `1, 1`, `2`  
Three steps - Four possible paths: `1, 1, 1`, `1, 2`, `2, 1`, `3`  
Four steps - Seven possible paths: `1, 1, 1, 1`, `1, 1, 2`, `1, 2, 1`, `1, 3`, `2, 1, 1`, `2, 2`, `3, 1`

You can see how this would be tough without recursion for larger numbers of steps. Top-down thinking helps.  
For 11 steps, the _subproblem_ is a 10-step staircase: Climbing 11 steps would take _at least_ as many steps as climbing a 10-step staircase.  
But we also know someone could jump to the top from stair 9 or 8 as well.  
Therefore: the number of steps to the top is at least the sum of all the paths to stairs 10, 9 and 8. Beyond this, you can't jump from stair 7 to 11.  
👉 `numberOfPaths(n - 1) + numberOfPaths(n - 2) numberOfPaths(n - 3)` is the core algorithm! 
#### Base Case
Base case for this problem is more difficult. We could hardcode it:
```go
func numberOfPaths(n int) int {
	switch {
	case n <= 0:
		return 0
	case n == 1:
		return 1
	case n == 2:
		return 2
	case n == 3:
		return 4
	default:
		return numberOfPaths(n -1) + numberOfPaths(n -2) + numberOfPaths(n - 3)
	}
}
```
🆒 But let's do better:
- We know that `numberOfPaths(1)` should return one. So `if n == 1 { return 1 }`.  
- We know `numberOfPaths(2)` should return 2, and it will compute as:
  - `numberOfPaths(1) + numberOfPaths(0) + numberOfPaths(-1)` 
  - Here, `numberOfPaths(1)` returns 1, so we tell `numberOfPaths(0)` to also return 1, which gives us the desired 2. ie:
  - `if n < 0; { return 0 }` `if n == 1 || n == 0; { return 1 }`
- We `numberOfPaths(3)` should return 4. it computes as:
  - `numberOfPaths(2) + numberOfPaths(1) + numberOfPaths(0)` which, per above if statement, will return 4.
```go
func numberOfPaths(n int) int {
	switch {
	case n < 0:             // base case
		return 0
	case n == 0 || n == 1:  // also base case
		return 1
	default:
		return numberOfPaths(n-1) + numberOfPaths(n-2) + numberOfPaths(n-3)
	}
}
```

### Anagram Generation
Anagrams are a reordering of all chars within a string. ex: "abc", anagrams are `"abc", "acb", "bac", "bca", "cab", "cba"`  
We could say that the subproblem of "abcd" is "abc". But how to use a func that gives us all anagrams of "abc" 
to produce all of "abcd?".  
```go
func anagramsOf(str string) []string {
	// Base case:
	if len(str) == 1 {
		return []string{string(str[0])}
	}

	var collection []string

	substrAnagrams := anagramsOf(str[1:]) // Find all anagrams of the substring.

	for _, subStrAnagram := range substrAnagrams { // Iterate over each substring.

		for i := 0; i < len(subStrAnagram)+1; i++ { // Iterate over each index of the substring.

			// Create a cpy of the substring, inserting the removed element once in each position
			cpy := subStrAnagram[:i] + string(str[0]) + subStrAnagram[i:]

			collection = append(collection, cpy)
		}
	}

	return collection
}
```
This introduces a new Big O Category. For three chars, a permutation starts with each of them, and each permutation picks
the middle char from one of the two remaining chars, and it's last char from the remaining chars. This is `3 * 2 * 1`, 6 permutations.  
For other string lengths:  
```
4 Chars: 4 * 3 * 2 * 1          anagrams
5 Chars: 5 * 4 * 3 * 2 * 1      anagrams
6 Chars: 6 * 5 * 4 * 3 * 2 * 1  anagrams
```
This is a **factorial** pattern.
Given N data elements, how many steps does the algo taks? For length of N, we create `N!` anagrams.  
This is **O(N!)**, aka _factorial time._ Very slow!

## Dynamic Programming
Dynamic programming is the process of optimizing recursve problems that have overlapping subproblems. There are two ways to do this:
_Memoization_ and "going bottom up". First, we'll look at some inefficient recursive funcs.
Recursion can sometimes cause excess time complexity. Example of a poorly constructed recursive func:  
**O(2<sup>N</sup>)**
```go
// max finds the greates number from an arrayunc max(slice []int) int {
func maxInefficient(slice []int) int {
	if len(slice) == 1 {
		return slice[0]
	}

	// Compare first element with greatest element of the remainder slice.
	if slice[0] > maxInefficient(slice[1:]) {
		return slice[0]
	}

	return maxInefficient(slice[1:])
}
```
Here, the recursive use of max in the if statement will cause an avalanche of recursive calls. Well break it down by analyzing the "bottom call".  
#### Max recursive walk-through
Given an array `[1, 2, 3, 4]`:  
`max([4])` ---up-the-call-chain---> `max(3, 4)`. Notice it again will call `max([4])` twice because `3` is not greater than the result `4`.  
This get's even worse when we move further up the call chain. `max([1, 2, 3, 4])` would call `max` a total of 15 times.  

To fix this, we store the result to a variable, and call max only once.  
**O(N)**
```go
func max(slice []int) int {
	if len(slice) == 1 {
		return slice[0]
	}

	maxOfRemainder := max(slice[1:])

	if slice[0] > maxOfRemainder {
		return slice[0]
	}

	return maxOfRemainder
}
```
Benchmarks:
```
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkMax-16                 80621228                13.39 ns/op
BenchmarkMaxInefficient-16       5628093               212.1 ns/op
```

### Overlapping Subproblems
Fibonacci adds sequence of numbers until infinity: `0, 1, 2, 3, 5, 8, 13, 21, 34, 55...`  
Example of inefficient **O(2<sup>N</sup>** fibonacci recursion:
```go
func fibInefficient(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return fibInefficient(n-2) + fibInefficient(n-1)
}
```
Because we _need_ to calculate both `fib(n - 2)` & `fib(n - 1)`, we can't just store one of the values. This is an example
of a _overlapping subproblem_, because both `fib(n - 2)` & `fib(n - 1)` will call many of the same funcs as one another.  
We can solve overlapping subproblems with _Memoization_.

### Memoization
Memoization reduces recursive calls by remembering previously computed functions.  
We'll use this to store `fib` results in a hash table. `fib(3)` would get stored with the result: `{3: 2}`.  
We do this by passing a hash table to the function.  
O(2N)-1, or **O(N)**
```go
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
```
**NOTE**: Using a closure is both more efficient and still allows input of a single integer. see [code](dynamicProgramming.go)

### Going Bottom Up
This is simply ditching recursion and using a different approach(i.e. loop) to solve the problem. This is parts of dynamic 
programming because it is still taking a problem that _could be_ solved recursively, and ensures duplicate calls aren't 
made for overlapping subproblems.  
**O(N)**
```go
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
```
If you run benchmarks comparing these solutions, `fibBottomUp` is the clear winner, dramatically.
```
BenchmarkFibInefficient20-16               34086             34563 ns/op
BenchmarkFibMemoization20-16              976226              1028 ns/op
BenchmarkFib20-16                        1000000              1107 ns/op
BenchmarkFibBottomUp20-16               100000000               11.48 ns/op
```
This demonstrates that bottom up is often better choice, unless recursion allows for more intuitive code and performance
is not a factor.

## Speedy Recursive Algorithms
_Quicksort_ is an efficient sorting algorithm that many programming languages implement, it's very efficient for average scenarios and
performs similarly to Insertion Sort and Selection sort in worse case scenarios. It relies on _partitioning_.

### Partitioning
To take a random value from an array, called a _pivot_, every `N < pivot` is placed to the left of the pivot, and `N > pivot`
is placed to the right. Ex:  
Given array `[0, 5, 2, 1, 6, 3]`, we choose the rightmost value, `3` as the pivot. We assign two pointers, one to the leftmost value
and one to the rightmost value(excluding the pivot). `0` and `6` are the two pointers in this example.
1. Left pointer moves right until a value >= pivot is reached.
2. Right pointer moves left until a value <= pivot is reached.
3. If left pointer has reached or gone beyond the right pointer, go to step 4. Otherwise, swap values that left & right pointers point to and repeat steps 1-3.
4. Swap the pivot with the value that the left pointer is currently pointing too.

#### Code Implementation: Partitioning
see [partitioning.go](partitioning.go)

### QuickSort
A combination of partitions and recursion. Conducts the following steps:
1. Partition the array.
2. Treat subarrays to left/right of the pivot as their own array & recursively repeat steps 1 & 2.
3. If subarray has 0 or 1 elements, this is the _basecase_.

#### Code Implementation: Quicksort
see [partitioning.go](partitioning.go)

### Quicksort Efficiency
First, determine efficiency of a _**single**_ partition:  
A partition has two primary steps:
* Comparisons: compare each of the values to the pivot. 
  * _At least_ N comparisons.
* Swaps: potentially swap values pointed to by left & right pointers
  * Depends on how input data is sorted. _At most_ N / 2 swaps per each partition.
  * We don't always swap. Randomly sorted input data would bring about N / 4 swaps.  

On average, this gives **N comparisons and N / 4 swaps**, about 1.25 steps per element. This is **O(N)** time.  
This is only a single partition, and quicksort may perform many.
### Quicksort Steps
Quicksort has a series of partitions, and each partition takes N steps for N elements of each subarray.  
For an array of 8 elements, quicksort takes about 21 steps. 

| **N** |  Quicksort steps (approx.)  |
|-------|:---------------------------:|
| 4     |              8              |
| 8     |             24              |
| 16    |             64              |
| 32    |             120             |
#### Big O of Quicksort
Looking above table, number of quicksort steps for N elements is about N * log N

| N   | log N | n * log N | Quicksort Steps (approx.) |
|-----|-------|-----------|---------------------------|
| 4   | 2     | 8         | 8                         |
| 8   | 3     | 24        | 24                        |
| 16  | 4     | 64        | 64                        |
| 32  | 5     | 160       | 160                       |

This is **O(N log N)**, because each time we partition an array, we break it down into two subarrays: there are log N 
halvings, and each having has a partition on all subarrays whose elements add up to N.  
NOTE: This is an approximation, as we must conduct **O(N)** partition on the original array as well.
#### Quicksort Worst Case Scenario
Best case: when pivot ends up in the middle of the subarray after partition, which generally occurs when the array values
are adequately mixed.  
The _worst case_ is when the pivot always ends up on one side, ie if the input is already a perfectly sorted array.  
**O(N<sup>2</sup>)** worst case time complexity.
#### Quicksort vs Insertion Sort

|                | Best Case  | Average Case    | Worst Case      |
|----------------|------------|-----------------|-----------------|
| Insertion Sort | O(N)       | O(N<sup>2</sup> | O(N<sup>2</sup> |
| Quicksort      | O(N log N) | O(N log N)      | O(N<sup>2</sup> |

### Quickselect
Allows us to find certain values in an unsorted array, _without sorting it_. I.e.: finding the tenth-lowest value in an
array, or the fifth highest, etc.  
Quickselect also uses partitioning. We can judge the location of the target value, based on the pivot index.  
**O(N)** for average scenarios. For N elements, we need `N + (N/2) + N/4) + (N/8) + ... 2` steps, which is roughly 2N steps.
#### Code Implementation: Quickselect
see [partitioning.go](partitioning.go)

### Sorting as Key to other Algos
Sorted arrays unlock possibilities to utilize other efficient algos. Ex: finding duplicates:

## Node-Based Data Structures
### Linked lists 
Linked lists are similar to arrays, but can be scattered throughout the computer's memory, via _Nodes_.
Each node contains some data and a _link_, the memory address of the next node in the list. A _head_ and a _tail_ are the first and last nodes in the list respectively.  
Initially, only the head node of a linked list is available for immediate access. Linked lists have four classic operations: reading searching, insertion and deletion.

#### Reading
Requires following the chain of the nodes' links.  
Worst case: O(N), must slower than array.
see [linkedlist.go](linkedList.go) for implementation.
#### Searching
Also O(N) search speed, much slower than array.
see [linkedlist.go](linkedList.go) for implementation.
#### Insertion
Here, linked lists have advantage over arrays _in some situations_.  
O(1) for insertion _at the beginning_. Worst case of O(N)+1 to insert at end of list. 
Create a new node, set it's `next node` to the current head, and change the linked list's `head` to the new node.  
Insertion _anywhere_ is just one step, but finding the node at a specific index is  
see [linkedlist.go](linkedList.go) for implementation.
#### Deletion
Speeds are the same as insertion. Deleting merely requires changing a node's link to point to the node that exists 
after the node targeted for deletion, or changing to nil if it's the tail.  
Note that the 'deleted' node still exists in memory, until garbage collected by GC.

### Linked List Efficiency

| Operation | Array               | Linked List               |
|-----------|---------------------|---------------------------|
| Reading   | O(1)                | O(N)                      |
| Search    | O(N)                | O(N)                      |
| Insertion | O(N) or O(1) at end | O(N) or O(1) at beginning |
| Deletion  | O(N) or O(1) at end | O(N) or O(1) at beginning |

The true power of a linked list is that the _actual insertion and deletion_ steps are just O(1). They're highly suitable for 
situations when an app will comb through existing data, making insertions/deletions when needed. This is better than an array,
which would have to move _all_ the data around to fit in a contiguous memory block.

### Doubly Linked Lists
Have two links: pointing to both the node before and after. This allows O(1) insertion time for the beginning and end of the linked list.  
This makes them highly suitable for a _queue_.  
See [doublyLinkedListQueue.go](doublyLinkedListQueue.go) for implementation.