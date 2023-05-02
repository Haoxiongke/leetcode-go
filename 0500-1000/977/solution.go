package main

func sortedSquares(nums []int) []int {
	l, r, index := 0, len(nums)-1, len(nums)-1
	result := make([]int, len(nums))
	for l <= r {
		num1 := nums[l] * nums[l]
		num2 := nums[r] * nums[r]
		if num1 < num2 {
			result[index] = num2
			r--
		} else {
			result[index] = num1
			l++
		}
		index--
	}
	return result
}
