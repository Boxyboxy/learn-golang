package main

// size is ecoded into the array's type e.g. [5]int
// []int -> this is known as a slice
// func Sum(numbers []int) int {
// 	result := 0
// 	for _, number := range numbers {
// 		result += number
// 	}
// 	return result
// }

//... ellipsis in a function parameter declaration is known as the variadic paramter syntax
// This function is a variadic function
// When calling sumAll, you can pass multiple []int arguments separated by commas
func SumAll(numbersToSum ...[]int) []int {
	// lengthOfNumbers := len(numbersToSum)
	// make alows the creation of slice with a starting capacity specified in the second parameter
	//sums := make([]int, lengthOfNumbers)
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// func SumAllTails(numbersToSum ...[]int) []int {
// 	// lengthOfNumbers := len(numbersToSum)
// 	// make alows the creation of slice with a starting capacity specified in the second parameter
// 	//sums := make([]int, lengthOfNumbers)
// 	var sums []int
// 	for _, numbers := range numbersToSum {

// 		if len(numbers) == 0 {
// 			sums = append(sums, 0)
// 		} else {

// 			sums = append(sums, Sum(numbers)-numbers[0])
// 		}
// 	}

// 	return sums
// }



func Sum(numbers []int) int {
	add:= func(acc, x int)int {return acc+x}
	return Reduce(numbers, add, 0)
}

// calcilates the sums of all but the first number given a collection of slices
func SumAllTails(numbersToSum ...[]int) []int {
	
	sumTail:= func(acc, x []int) []int {
		if(len(x)==0) {
		return append (acc, 0)
		} else{
			tail:=x[1:]
			return append(acc, Sum(tail))
		}
	}
	return Reduce(numbersToSum, sumTail, []int{})
	
}