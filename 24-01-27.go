package main

//https://leetcode.com/problems/count-elements-with-maximum-frequency/
func maxFrequencyElements(nums []int) (res int) {
	type frequency map[int]int
	freq := frequency{}
	for _, v := range nums {
		freq[v] += 1
	}
	maxV := 0
	for _, v := range freq {
		if v >= maxV {
			maxV = v
		}
	}
	for _, v := range freq {
		if v == maxV {
			res += v
		}
	}
	return
}
