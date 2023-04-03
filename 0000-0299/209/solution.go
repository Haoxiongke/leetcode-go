package main

func minSubArrayLen(target int, nums []int) int {
	len := len(nums)
	l := 0
	result := len + 1 // 初始化返回len+1,防止没有满足target的子序列
	var sum int
	for r := 0; r <= len-1; r++ {
		sum += nums[r]
		for sum >= target {
			if result > r-l+1 {
				result = r - l + 1
			}
			sum -= nums[l]
			l++
		}
	}

	// result 如果没有没赋值过，则返回0
	if result == len+1 {
		return 0
	} else {
		return result
	}
}

func main() {
	println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
}
