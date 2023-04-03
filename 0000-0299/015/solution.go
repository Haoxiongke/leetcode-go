package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		num := 0 - nums[i]
		for l < r {
			if num == nums[l]+nums[r] {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if num > nums[l]+nums[r] {
				l++
			} else {
				r--
			}
		}
	}
	return result
}

func main() {
	nums := []int{1, -1, -1, 0}
	i := threeSum(nums)
	print(i)
}
