package main

//https://leetcode.com/problems/count-elements-with-maximum-frequency/
func maxFrequencyElements(nums []int) (res int) {
	type frequency map[int]int
	freq := frequency{}
	for _, v := range nums {
		freq[v] += 1
	}
	maxV, maxK := 0, 0
	for k, v := range freq {
		if v >= maxV && k >= maxK {
			maxK, maxV = k, v
			res = k * v
		}
	}
	return
}
