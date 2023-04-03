package main

func isAnagram(s string, t string) bool {
	record := [26]int{}

	for _, r := range s {
		record[r-'a']++
	}

	for _, r := range t {
		record[r-'a']--
	}
	return record == [26]int{}
}

func main() {
	a := [26]int{}
	println(a['b'-'a'])
}
