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

## Bubble Sort - A quadratic algo
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
We'll use Selection Sort, and compare it to Bubble Sort.  
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