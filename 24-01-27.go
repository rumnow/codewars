package main

import (
	"slices"
)

//https://leetcode.com/problems/count-elements-with-maximum-frequency/
func maxFrequencyElements(nums []int) (res int) {
	slices.Sort(nums)
	seq := false
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			if seq {
				res++
			} else {
				res+=2
			}
			seq = true
		} else {
			seq = false
		}
	}
	return
}
