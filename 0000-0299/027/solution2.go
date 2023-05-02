package main

func removeElement2(nums []int, val int) int {
	s, f := 0, 0
	for ; f < len(nums); f++ {
		if nums[f] != val {
			s++
		}
	}
	return s
}
