package main

func partition(s string) [][]string {
	var result [][]string
	var path []string
	var backtracking func(index int)

	backtracking = func(index int) {
		// 到达最后的长度
		if index == len(s) {
			tmp := make([]string, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}

		for i := index; i < len(s); i++ {
			// 切割
			str := s[index : i+1]
			// 剪枝
			if !isPalindrome(str) {
				continue
			}
			path = append(path, str)
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtracking(0)

	return result
}

func isPalindrome(str string) bool {
	l, r := 0, len(str)-1
	for l < r {
		if str[l] != str[r] {
			return false
		}
		l++
		r--
	}
	return true
}
