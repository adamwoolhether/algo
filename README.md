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

## Binary Search Trees
Binary search trees maintain ordering and have fast search, insertion, and deletion speeds.

### Trees
A ndoe-based data structure, where each node can have links to multiple nodes. There are many kinds of tree-based data structures.
* Root: the uppermost node.
* Parent/children: The node immediately above/nodes immediately below.
* Descendants/ancestors: all the nodes above/below a given node.
* Levels: A row within a tree.
* A tree is _balanced_ when it's nodes' subtrees have the same number of nodes in it, otherwise it is imbalanced.

### Binary Search Trees
_Binary tree_: a tree where each node has _zero, one, or two_ children.
_Binary search tree_: a binary tree that has the following additional conditions:
* Each node has at most one left child and one right child.
* A node's left descendants contains values less than the node itself, and its right descendants contain values only greater than itself.    
see [binarySearchTree.go](binarySearchTree.go) for node implementation.

### Searching
Searching within a binary search tree algorithm:
1. Designate a node as "current node". (Initially, the root.)
2. Inspect the value of current node.
3. If found, return.
4. If value is less than current node, search in its left subtree.
5. if value is greater than current node, search in its right subtree.
6. Repeat 1-5 until value is found or bottom of tree is reached, indicating the value isn't in the tree.
#### Efficiency
**O(log N)** Time Complexity in best-case scenario, for a _perfectly_ balanced search tree.  
👉**Log(N) Levels** _If there are N nodes in a balanced binary tree, there will be about log N levels(rows)._  
Code implmentation: [binarySearchTree.go](binarySearchTree.go)

### Insertion
**O(log N) + 1** steps  
Insertion always takes 1 extra step beyond a search. In contrast, an ordered array takes O(N).  
Code implmentation: [binarySearchTree.go](binarySearchTree.go)
#### Order of Insertion
Well-balanced trees are only created when input randomly sorted data. Sorted data can result in an imbalanced tree.  
If it was completely-sorted, the tree would have all data on the right child, resulting in a completely linear O(N) search.  
⚠️ It's best to first randomize data before inserting it into a tree.

### Deletion
**O(log N)**
Deletion is the most complex operation of a binary search tree. The proper action depends on whether the node has a single, two, or no child nodes.  
The complete algorith for deletion:
* If the node being deleted has no children, simply delete it.
* If the node being deleted has one child, delete the node and replace it with the child.
If the node ebing deleted has two children, replace the node with the _successor_ node. 
  * The _successor_ is the child node whose value is _the least of all values that are greater than the deleted node._
  * The find the successor: visit the right child of the deleted value, then keep visiting the left child of each subsequent child
until there are no more left children. This bottom value is the successor node.
  * If the successor node has a right child, after plugging the successor node into the deleted node's spot, take the former 
successor node's right child and turn it into the _left child of the successor node's former parent_.   

Code implmentation: [binarySearchTree.go](binarySearchTree.go)

### Binary Search Tree Traversal
_Traversing_ is the act of visiting every node in a data structure. There are many ways to traverse a tree,
we'll implement what is known as _inorder traversal_. The `traverse` method should have the following steps:
1. Call itself recursively on the node's left child until there is no left child.
2. "Visit" the node. Here, we'll print the title.
3. Recursively call it's on the node's right child until there is no right child.

## Heaps
There are many types of tree data-structures, it's important to use the proper one for a given use case.

### Priority Queues
Priority queues are a list whose deletions and accesses are the same as a classic queue(front), but insertions work more
like an ordered array. Data remains in sorted order on insertions. (Think of a hospital room, where patients in more
critical condition will be seen first).  
As an abstract data type, the underlying data structure is flexible. Using an array allows us to:
* Ensure proper ordering on inserts
* Data is removed from the end of the array(the 'front' of the array).

Two primary operations:
* Deletions; If data is array-based: O(1)
* Insertions; If data is array-based: O(N)    
There is another, more efficient data structure for priority queues than a sorted array...

### Heaps
Many types, we'll focus on a _binary heap_, a kind of binary tree. Binary heaps also come in two varieties: _max-heap_ & _min-heap_.  
Examples here use a max-heap, but the implementing the other is a small difference.  
Max-heaps have two conditions:
* Each node's value must be greater than each of its descendants, this is the _heap condition_.
  * Contrast this is binary search tree, where the right child will be greater than the node. "A binary serach tree doth not a heap make."
  * A _min-heap_ would have the opposite condition: the node's value would be greater than all of its descendants.
* The tree must be _complete_, meaning if must be completely filled with nodes.
  * The bottom row _can_ have empty positions as no nodes are to the right of these empty positions.

### Heap Properties
* Heap-ordering is useless regarding searches, because we wouldn't know to look on the left or right of a node for a given value.  
So, heaps are said to be _weekly ordered_, compared to binary search trees. They have _some_ order(nodes greater than both children),
but not enough to improve time-complexity of searching. 
* Root nodes in a max-heap will always have the greatest value (or the least, in a min-heap).
* Twp primary ops: inserting and deleting. Optional 'read' looks at value of root node.
* A _last node_ is the heap's rightmost node in the bottom level.

### Heap Insertion
Heap insertion involved the following algorithm: O(log N)
1. Create new node with given value and insert it as the new _last node_. (see below)
2. Compare the new node with its parent.
3. If new node is greater than the parent, swap it with the parent.
4. Repeat step 3, moving the node up the head until it's parent has a greater value, this process is called _trickling_.
#### Looking for the Last Node
The first step in the heap insertion algorithm requires that we find the last node. This is the "Problem of the Last Node" (see below)

### Heap Deletion
Only the root node is ever deleted, just like a priority queue.
Algorithm for heap deletions: O(log N)
1. Remove the root node by moving the last node where the root node was.
2. Trickle the root node down to its proper place. This is more complex than trickling up.
   * Check both children of the "trickle node" to see which is larger.
   * If trickle node is smaller than larger of the two child nodes, swap the trickle not with that larger node. (Otherwise, the heap condition would be violated)
   * Repeat steps 1 & 2 until the trickle node has no children greater than it.

### Heaps vs Ordered Ararys
|           | Ordered Array  | Heap      |
|-----------|----------------|-----------|
| Insertion | Slow           | Very fast |
| Deletion  | Extremely fast | Very fast |
Even though the heap is slightly slower at deletions, it performs very well in all, so it's the better choice for priority queues.

### Problem of the Last Node
This addresses the problem of how we find the actual last node in a heap, for insertion & deletions. Using another node
would result in a heap that is incomplete. Completeness ensures that the heap remains _well-balanced_, and being well-balanced
is what helps us achieve O(log N) operations.  
So, what algorithm helps us quickly find the last node of a heap? Well, we use arrays!

### Arrays as Heaps
To solve the "Problem of the last node", heaps are usually implemented using arrays.
```
             0
           /   \
          1     2
        /  \  /   \ 
       3    4 5    6
```
The figure above demonstrates has the elements of an array can be represented as a heap.  
With this implementation, _the last node in the heap will always be the first element of the array_, at index 0.    
see [heap.go](heap.go) for implementation.

#### Traversing an Array-Based Heap
Moving node-to-node would be simple with a linked list, how to do so with an array?  
Assigning index's of the heap's nodes, as shown above, means the traits of a heap are always true:
* To find the left child, use this formula: `(index * 2) + 1`
* To find the right child, use this formula: `(index * 2) + 2`
* To find a node's parent, use this formula: `(index - 1) / 2`
Go ahead and try it out with the above example!

### Bonus: Heapsort
Heapsort is a sorting algorithm in which all values are inserted into the heap, and then popped out. A max-heap would
return values in _descending order_, and a min-heap would return values in _ascending order_.  
This would give us a sort algorithm in O(N log N).

## Tries
Tries: A type of tree, great for applications that deal with text/numbers, enabling features like autocorrect and autocomplete.  
Name comes from the word _retrieval_, but is pronounced as "try". Tries aren't as well documented and have many implementations.

#### Trie Nodes
Unlike binary trees, trie nodes can have _any_ amount of child nodes. We'll implement a trie nodes that contain a hash table, with 
English letter keys, and values that point to other nodes.

### Storing Words
Ex: store three words: `"ace", "bad", "cat"` would create a node for each char in each word:
```
  {"a": ptr "b": ptr "c": ptr}
        /         |        \
 {"c": ptr}  {"a": ptr}  {"a": ptr}
       |            |           |
 {"e": ptr}  {"d": ptr}  {"t": ptr}
       |            |           |
 {"*": ptr}  {"*": ptr}  {"*": ptr}
```
The `"*"` indicates that there is a complete word that ends here. Adding the word "act" would look like:
```
  {"a": ptr "b": ptr "c": ptr}
        /         |        \
 {"c": ptr}  {"a": ptr}  {"a": ptr}
       |            |_________  |_________
       |                      |           |
 {"e": ptr, "t": ptr}  {"d": ptr}  {"t": ptr}
       |          |           |           |
 {"*": ptr} {"*": ptr}  {"*": ptr}  {"*": ptr}
```
If a word like "batter" was added, the node being pointed by the 't' in 'b-a-t' would look like: `{"*": nil, "t": ptr}`.

### Trie Search
The classic trie operation. Two types:
* Determine whether the string is a _complete_ word
* Determine whether the string is at least a a word _prefix_.(We implement this here).

The prefix-search algorithm uses the following algorithm:
1. Declare a `currentNode` var, initally assigning the root node
2. Iterate over each char of the string
3. Check if `currentNode` has a child with the same key as the current char
4. If not, return `nil`
5. Otherwise, update the `currentNode` as that child. Repeat from step 2
6. If the end of the string is reach, the search string has been found.

#### Efficiency of Trie Search
The algo takes _as many steps as there are characters in the search string._ However, it's not quite O(N), because N refers
to the amount of data in the data structure (the number of nodes in the trie, which is usually much greater than the number of
chars in the string.  
#### O(K) 
O(K) is usually used to describe the time complexity of Tries. O(K) isn't constant: the string size can vary.  
But, it has one important similarity to constant time: the trie itself can grow tremendously, but it won't affect the search speed.
In other words, a trie search on a string with 3 chars will always take three steps. Contrast this with non-constant algos, which are tied to the amount of data in the ds.  
This means O(K) is extremely efficient.

### Trie Insertion
Algo is similar to search:
1. Declare a `currentNode` var, initally assigning the root node
2. Iterate over each char of the string
3. Check if `currentNode` has a child with the same key as the current char
4. If so, update `currentNode` to become that child node & go back to step 2, moving on the next char.
5. Otherwise, create the child node with the char and update `currentNode` to point to this new node.
6. After inserting the final char, add a `"*"` child to the last node.
#### Speed
Also **O(K)**. Technically O(K+1), for the added "*" at the end.

### Building Autocomplete
This method will return an array of all the words in the trie, allowing us to start from any node in the trie.
see [trie.go](trie.go)

### Tries With Values: Improving Autocomplete
To display words that are more popular than others, a 'popularity rating' can be stored in the trie as well. In this implementation,
the `"*"` is perfect for this. Naturally, this would require refactoring of the acceptable value type in the underlying map.

## Graphs
Graphs are data structures that specialize in relationships, easily expressing how data is connected. Ex: In social networks, users are represented by _nodes_, with connecting lines representing a friendship.

#### Graphs vs Trees
Trees are a type of graph.  
👉All trees are graphs, but not all graphs are trees.  
A graph must not have _cycles_ to be considered a tree, and all nodes must be connected.  
A _cycle_ is when nodes reference each other circularly.  
Additionally, all nodes in a tree must be connected, but a graph allows nodes that may not be connected.

#### Graph Jargon
* _Vertex_: A node in a graph.
* _Edges_, aka _vertices_: The lines between nodes.
* _Adjacent_ describes nodes that are connected to each other. Tese are known as _neighbours_.
* _Path_ a specific sequence of edges from one vertex to another.

#### Bare-Bones Graphs
As a simple example, a hash map can be used to implement a simple graph. Here is one conveying a simpel social network:
```go
friends := map[string][]string{
    "Alice": []string{"Bob", "Diana", "Fred"},
    "Bob": []string{"Alice", "Cynthia", "Diana"},
    "Cynthia": []string{"Bob"},
    "Diana": []string{"Alice", "Bob", "Fred"},
    "Elise": []string{"Fred"},
    "Fred": []string{"Alice", "Diana", "Elise"},
}
```

### Directed Graphs
A social network may allow _non-mutual_ relationships: Alice can follow Bob, but Bob might not follow Alice.
```
   --------ALICE-------
   |                  |
   v                  v 
CYNTHIA -----------> BOB
    ^----------------|
```
Arrows indicate the _direction_ of the relationship. Above, Alice follows Bob and Cynthia and no one follows Alice. Bob and Cynthia follow each other.  
Expressed as a hashmap:
```go
followees := make map[string][]string{
        "Alice": []string{"Bob", "Cynthia"},
		"Bob": []string{"Cynthia"},
		"Cynthia": []string{"Bob"}
    }
}
```

### Object-Oriented Graph Implementation
see [graph.go](graph.go)  
Our implementation uses a slice to store neighbors. This is known as the _adjacency list_ implementation.  
Another implementation, known as _adjacency matrix_ uses two-dimenstional arrays instead of a list.

We'll use a directed graph in our implementation. A directed graph would add vertex to a vertex's list of neighbors as so:
```go
func (v *vertex) AddNeighbor(vertex vertex) {
	v.neighbors = append(v.neighbors, vertex)
}
```
Whereas an undirected implementation would mutually add a node to their respective list of neighbors:
```go
// AddNeighborUndirected appends a given vertex to the calling vertex's list of neighbors.
func (v *vertex[T]) AddNeighborUndirected(vertex *vertex[T]) {
	// Prevent an infinite loop.
	for _, n := range v.neighbors {
		if n == vertex {
			return
		}
	}

	v.neighbors = append(v.neighbors, vertex)
	v.AddNeighborUndirected(v)
}
```

### Graph Search
Searching for a vertex is a common graph operation, and has more specific meaning.  
👉 _If we have access to one vertex in the graph, we must find another particular vertex that is somehow connected to this vertex._  
This is so because there may be more than one path from one vertex to another.  
Searching in graphs has many use-cases:
1. Searching for a particular vertex within a connected graph. This can be done to find _any_ vertex in the graph if we have access to just one vertex.
2. Discovering whether two vertices are connected.
3. To merely traverse the graph, alllowing use to perform an operation on every vertex.  
Two well known approaches go graph searches: _depth-first search_ & _breadth-first search_.

### Depth-First Search
aka DFS. Very similar to Binary Tree Traversal algo.  
🔑 Key to graph search algos is keeping track of which vertices have been visited so far. This is how we prevent endless cycles.
One way to do so is by using a hash table.  
Algo:  
1. Start at any random vertex within the graph.
2. Add current vertex to hash table, marked as visited.
3. Iterate through the vertex's neighbors.
4. Ignore neighbors that have been visited already.
5. If the neighbor hasn't been visited, recursively perform a depth-first search on that vertex.

### Breadth-First Search
BFS doesn't use recursion, instead, it uses a queue.  
Algo:
1. Start from any vertex, the "starting vertex".
2. Add it to the hash table, marked as visited.
3. Add starting vertex to a queue.
4. Start a loop that runs as long as the queue isn't empty.
5. In the loop, remove the first vertex from the queue, the "current vertex".
6. Iterate over all the neighbors of the current vertex.
7. Ignore neighbors that have been visited already.
8. If it hasn't been visited, mark it as visited and add it to the queue.
9. Repeat loop (step 4) until queue is empty.

### DFS vs. BFS
Breadth-first will go through all the calling vertex's immediate connections before spiralling out and moving
further from the caller. In contrast, depth-first immediately moves as far away as possible from the calling vertex
until being forced to return to it.  
The correct choice depends on the usecase. Finding all _direct_ connections would be perfect for breadth-first. Finding a specific
grandchild in a family tree would be served well with a depth-first search. Always ask:  
⁉️ _Do we want to stay close to the starting vertex, or specifically move far away?_ Breadth-first is good for staying close, and depth first is good for moving away quickly.

### Efficiency of Graph Search
Because we touch all vertices in the graph, the speed seems to be O(N), where N is the number of vertices.  
But, we have to also iterate over all neighbors for each vertex that is traversed.  
As an example: A graph has the following 5 vertices:
A, B, C, D, E. A has four neighbors, and the rest each have three. This would bring a total of 16 iterations + 5 visiting of each = 21 steps.  
Another graph has 5 vertices:  V, W, X, Y, Z. V has four neighbors, the rest have only one. This brings a total of 8 iterations + 5 visiting of each = 13 steps.  
👉Determining the Big O requires that we count how many vertices are in the graph AND how many neighbors each vertex has. We'd
need _two_ variables to describe the time complexity. This is `O(V + E)`

### O(V + E)
_V_: _vertex_, representing the number of vertices in the graph.  
_E_: _edge_, the number of edges in the graph.  
The number of steps is the number of vertices plus the number of edges. Each edge is actually touched twice, but because constants are dropped in Big O, `V + 2E` is simplified to `O(V+E)`.  
That said, the choice of BFS or DFS can help optimize searches in regard to the use case.

### Weighted Graphs
These add additional information to edges of the graph, a _weight_.  
To use them in code, our `neighbors` array will be refactored to a hash table, with the key holding the vertex and the weight as the value.

#### The Shortest Path Problem
Weighted graphs help when modeling many types of dataset problems. One of these is known as the **Shorted Path Problem**.  
If we have a graph of cities with airplane ticket prices, how to create an algorithm that finds the cheapest price to get from one city to the next?
Here, the weight is a price, but it can also be a distance.

## Dijkstra's Algorithm
One of the most famous algos that can solve the shorted path problem, created by Edsger Dijjkstra in 1959.
Using the airplane ticket price example below, we'll apply Dijkstra's algo to solve it. It has the added bonus of not only finding the cheapest price
from our current city to destination city, but from our current city to _all_ known cities.
#### Setup
Create a way to store chepest known prices from starting city to all other known destinations.  
`Cheapest Prices Table`

| From ATL to: | City #1 | City #2 | City #3 | etc |
|--------------|---------|---------|---------|-----|
|              | $       | $       | $       | $   |
We also need another table, the `Cheapest Previous Stopover City Table`

| Cheapest Previous Sopover City from ATL | Boston | Chicago | Denver | El Paso |
|-----------------------------------------|--------|---------|--------|---------|
|                                         | city   | city    | city   | city    |

#### Dijkstra's Algorithm Steps
Here, "city" is synonymous with "vertex".
1. Visit starting city, making it the "current city".
2. Check prices from current city to its neighbors.
3. If price to a neighbor from starting city is cheaper than price currently in `cheapest prices table`:
   1. Update the `cheapest prices table` to reflect this cheaper price.
   2. Update the `cheapest previous stopover city table`, making the neighbor city the key and the current city its value.
4. Then visit the unvisited city that has the cheapest price from the starting city, making it the current city.
5. Repeat steps 2 to 4 until every known city has been visited.  
See pages 369-377 for a more detailed walk through of the algorithm.

#### Efficiency of Dijkstra's Algorithm
Dijsktra's algo has many variations, as it's a general description of the approach to find the shortest path within a weighted graph.  
(NOTE: Remember to implement use of a priority queue instead of a slice).  
The initial implementation, using a slice, has a speed of O(V<sup>2</sup>).

## Space Constraints
Space constraints measure how much memory an algo consumes.

### Big O of Space Complexity
The key question in regards to space complexity:  
🔑 _If there are N data elements, how many units of memory will the algorithm consume?_  
Describing space complexity with Big O only counts _new_ data that is being generated, aka _auxiliary space_.  
Some references will include the original space being used, be aware of this. Here, we do not.

An example of an algo with space complexity O(N), because it generates an additional N data elements:
```go
func makeUppercaseInefficient(words []string) []string {
	newSlice := make([]string, len(words))
	for i, word := range words {
		newSlice[i] = strings.ToUpper(word)
	}

	return newSlice
}
```
Let's make it more efficient and not consume additional memory with O(1):
```go
func makeUppercase(words []string) []string {
	for i, word := range words {
		words[i] = strings.ToUpper(word)
	}
	
	return words
}
```

### Tradeoffs Between Time and Space
Whether one of two algorithms is "more" efficient may depend on what matters to you most: time or space. Remember the func on line 49 that finds duplicate values:  
Time Complexity: O(N<sup>2</sup>), Space Complexity: O(1)
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
and consider this version, which uses a hash map but is faster:  
Time Complexity: O(N), Space Complexity: O(N)
```go
func hasDuplicateValueFaster(arr []int) bool {
	existingVals := make(map[int]bool, len(arr))
	for _, v := range arr {
		if !existingVals[v] {
			existingVals[v] = true
		} else {
			return true
		}
	}
	
	return false
}
```
A third version, we've seen before, strikes a good balance between the two:
Time Complexity: O(N log N), Space Complexity: O(log N)
```go
func hasDuplicateValue[T constraints.Ordered](input []T) bool {
	sort.Slice(input, func(i, j int) bool { return input[i] < input[j] })

	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			return true
		}
	}

	return false
}
```

### The Hidden Cost of Recursion
Even if our algo doesn't make new arrays or hash maps, it's possible we might be consuming more space, as recursive functions 
keep adding to the call stack.  
If each function takes up O(N) space, than a func that makes 100 recusive calls will need enough memory to store 100 func calls in the call stack.  
👉 _Recursive functions take up a unit of space for each recursive call it makes._  Properly calculating this requires determining how large the call stack
would be at its peak.  
As an example, quicksort make O(log N) recursive calls, so it has a call stack the size of log(N) at its peak.

The `wordBuilder` algo on line 155 has a Space Complexity of O(N<sup>2</sup>)