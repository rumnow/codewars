package main

import (
	"fmt"
	"math"
)

//The colors used by the printer are recorded in a control string. For example a "good" control string would be aaabbbbhaijjjm meaning that the printer used three times color a, four times color b, one time color h then one time color a...
// Sometimes there are problems: lack of colors, technical malfunction and a "bad" control string is produced e.g. aaaxbbbbyyhwawiwjjjwwm with letters not from a to m.
// You have to write a function printer_error which given a string will return the error rate of the printer as a string representing a rational whose numerator is the number of errors and the denominator the length of the control string. Don't reduce this fraction to a simpler expression.

func PrinterError(s string) string {
	denom := len(s)
	numer := 0
	for _, v := range s {
		if v > 'm' || v < 'a' {
			numer++
		}
	}
	return fmt.Sprintf("%v/%v", numer, denom)
}

// Complete the function that takes a non-negative integer n as input, and returns a list of all the powers of 2 with the exponent ranging from 0 to n ( inclusive ).
func PowersOfTwo(n int) []uint64 {
	res := make([]uint64, n + 1)
	for i := 0; i <= n; i++ {
		res[i] = uint64(math.Pow(2, float64(i)))
	}
	return res
}

//Create a combat function that takes the player's current health and the amount of damage recieved, and returns the player's new health. Health can't be less than 0.
func combat(health, damage float64) float64 {
	result := health - damage
	if damage > health {
		result = 0
	}
	return result
	//return max(health - damage, 0)
	//return math.Max(health - damage, 0)
}

//Given a set of numbers, return the additive inverse of each. Each positive becomes negatives, and the negatives become positives.
func Invert(arr []int) []int {
	res := make([]int, len(arr))
	for i, v := range arr {
		res[i] = v * -1
	}
	return res
}
