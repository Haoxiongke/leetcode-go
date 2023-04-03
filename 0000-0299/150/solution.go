package main

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, token := range tokens {
		num, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, num)
		} else {
			num1 := stack[len(stack)-2]
			num2 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, num1/num2)
			}
		}
	}

	return stack[0]
}

func evalRPN1(tokens []string) int {
	stack := make([]string, 0)

	for i := 0; i < len(tokens); i++ {
		if tokens[i] != "+" && tokens[i] != "-" && tokens[i] != "*" && tokens[i] != "/" {
			stack = append(stack, tokens[i])
		} else {
			num1, _ := strconv.Atoi(stack[len(stack)-2])
			num2, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-2]
			switch tokens[i] {
			case "+":
				stack = append(stack, strconv.Itoa(num1+num2))
			case "-":
				stack = append(stack, strconv.Itoa(num1-num2))
			case "*":
				stack = append(stack, strconv.Itoa(num1*num2))
			case "/":
				stack = append(stack, strconv.Itoa(num1/num2))
			}
		}
	}
	res, _ := strconv.Atoi(stack[0])
	return res
}

func main() {

}
