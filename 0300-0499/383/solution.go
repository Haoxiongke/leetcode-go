package main

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[rune]int, 1)
	for _, i := range ransomNote {
		m[i-'a']++
	}

	for _, i := range magazine {
		m[i-'a']--
	}

	for _, value := range m {
		if value > 0 {
			return false
		}
	}
	return true
}

func canConstruct2(ransomNote string, magazine string) bool {
	record := [26]int{}

	for _, m := range magazine {
		record[m-'a']++
	}

	for _, r := range ransomNote {
		record[r-'a']--
		if record[r-'a'] < 0 {
			return false
		}
	}

	return true
}

func main() {

}
