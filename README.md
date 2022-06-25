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
| Big O            |  Pronounciation  | Time Complexity | Notes                   |
|:-----------------|:----------------:|:----------------|:------------------------|
| O(1)             |    "O of one"    | Constant time   |                         |
| O(log N)         |   "O of log N"   | Log time        |                         |
| O(N)             |     "O of N"     | Linear time     | O(log<sub>2</sub> N)    |
| O(N<sup>2</sup>) | "O of N squared" | Quadratic time  | Typical of nested loops |

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