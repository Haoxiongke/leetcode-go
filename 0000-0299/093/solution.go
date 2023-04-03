package main

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	var result []string
	var path []string
	var backtracking func(index int)

	backtracking = func(index int) {
		// 只要四段就不在递归了
		if len(path) == 4 {
			if len(s) == index {
				result = append(result, strings.Join(path, "."))
			}
			return
		}

		for i := index; i < len(s); i++ {
			str := s[index : i+1]
			if !isSuitIp(str) {
				break
			}

			path = append(path, str)
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtracking(0)

	return result
}

func isSuitIp(s string) bool {
	if len(s) < 1 {
		return false
	}
	if len(s) > 1 && s[0] == '0' {
		return false
	}
	i, _ := strconv.Atoi(s)
	if i >= 0 && i <= 255 {
		return true
	}
	return false
}
