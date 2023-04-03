package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l, r := j+1, len(nums)-1
			num := target - nums[i] - nums[j]
			for l < r {
				if num == nums[l]+nums[r] {
					result = append(result, []int{nums[i], nums[j], nums[l], nums[r]})
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
	}
	return result
}

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	result := fourSum(nums, 0)
	fmt.Println(result)
}
