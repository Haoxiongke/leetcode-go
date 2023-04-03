package main

func reverseStr(s string, k int) string {
	ss := []byte(s)
	length := len(s)
	for i := 0; i < length; i += 2 * k {
		if i+k < length {
			reverseString(ss[i : i+k])
		} else {
			reverseString(ss[i:length])
		}
	}
	return string(ss)
}

func reverseString(b []byte) {
	l, r := 0, len(b)-1
	for l < r {
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}
}

func main() {
	println(reverseStr("ab", 2))
}
