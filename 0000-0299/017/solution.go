package main

func letterCombinations(digits string) []string {
	var result []string

	if digits == "" {
		return result
	}

	// 存放结果
	var str string

	// map
	m := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	var backtracking func(index int)

	backtracking = func(index int) {
		if len(str) == len(digits) {
			result = append(result, str)
			return
		}

		for _, letter := range m[digits[index]] {
			str += letter
			backtracking(index + 1)
			str = str[:len(str)-1]
		}
	}

	backtracking(0)

	return result
}
