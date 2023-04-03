package main

func isHappy(n int) bool {
	set := make(map[int]bool, 0)
	for 1 != n && !set[n] {
		n, set[n] = getSum(n), true
	}
	return 1 == n
}

func getSum(n int) (sum int) {
	sum = 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return
}

func main() {

}
