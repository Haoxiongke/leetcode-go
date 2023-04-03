package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	var count int

	m := make(map[int]int, 0)

	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			//if _, ok := m[num1+num2]; ok {
			//	m[num1+num2]++
			//} else {
			//	m[num1+num2] = 1
			//}
			// go map 不存在默认值为0
			m[num1+num2]++
		}
	}

	for _, num3 := range nums3 {
		for _, num4 := range nums4 {
			//if _, ok := m[0-num3-num4]; ok {
			//	count += m[0-num3-num4]
			//}
			count += m[-num3-num4]
		}
	}

	return count
}

func main() {

}
