package main

// In this simple exercise, you will build a program that takes a value, integer , and returns a list of its multiples up to another value, limit . If limit is a multiple of integer, it should be included as well. There will only ever be positive integers passed into the function, not consisting of 0. The limit will always be higher than the base.
// For example, if the parameters passed are (2, 6), the function should return [2, 4, 6] as 2, 4, and 6 are the multiples of 2 up to 6.
func FindMultiples(integer, limit int) []int {
	// Your code here!
	result := []int{}
	for i := integer; i <= limit; i += integer {
		result = append(result, i)
	}
	return result
  }

// I have a cat and a dog.
// I got them at the same time as kitten/puppy. That was humanYears years ago.
// Return their respective ages now as [humanYears,catYears,dogYears]
func CalculateYears(years int) (result [3]int) {
	// Write your solution here
	catYears := 15
	dogYears := 15
	for i := 2; i <= years; i++ {
		if i == 2 {
			catYears += 9
			dogYears += 9
		} else {
			catYears += 4
			dogYears += 5
		}
	}
	return [3]int{years, catYears, dogYears}
	// switch years {
	// 	case 1: result = [3]int{1, 15, 15}
	// 	case 2: result = [3]int{2, 24, 24}
	// 	default: result = [3]int{years, 24 + 4 * (years - 2), 24 + 5 * (years - 2)}
	// }
}

// Given two integers a and b, which can be positive or negative, find the sum of all the integers between and including them and return it. If the two numbers are equal return a or b.
// Note: a and b are not ordered!
func GetSum(a, b int) int {
	mi := func(a, b int) int {
		if a < b {return a} else {return b}
	}
	ma := func(a, b int) int {
		if a > b {return a} else {return b}
	}
	res := 0
	for i := mi(a, b); i <= ma(a, b); i++ {
		res += i
	}
	return res
	// if a > b {
	// 	a, b = b, a
	//   }
	//   return (a + b) * (b - a + 1) / 2
}
