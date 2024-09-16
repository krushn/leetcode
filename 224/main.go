package main

import (
	"fmt"
)

func main() {
	s := "(1+(4+5+2)-3)+(6+8)" //1-(     -2) // 2-1 + 2 //2147483647 //(1)
	fmt.Println(calculate(s))
}

func calculate(s string) int {
	sum := 0
	sign := 1
	var stack []int
	for i := 0; i < len(s); i++ {
		switch {
		case s[i] >= '0' && s[i] <= '9':
			num := 0
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			sum += num * sign
			i--
		case s[i] == '+':
			sign = 1
		case s[i] == '-':
			sign = -1
		case s[i] == '(':
			stack = append(stack, sum, sign)
			sum = 0
			sign = 1
		case s[i] == ')':
			sum = sum * stack[len(stack)-1]
			sum += stack[len(stack)-2]
			stack = stack[:len(stack)-2]
		}
	}

	return sum
}

/*

	var stack []string

	result := 0

	s = strings.ReplaceAll(s, " ", "")

	if strings.Index(s, "+") == -1 && strings.Index(s, "-") == -1 {

		s = strings.ReplaceAll(s, "(", "")
		s = strings.ReplaceAll(s, ")", "")

		result, _ = strconv.Atoi(s)
		return result
	}

	for i := 0; i < len(s); i++ {

		if string(s[i]) != ")" {
			stack = append(stack, string(s[i]))

			if i == len(s)-1 {
				result, stack = calculateStack(stack, result)
			}
		} else {
			result, stack = calculateStack(stack, result)
		}
	}

	return result

func calculateStack(stack []string, result int) (int, []string) {
	//calculate until we get "("

	lastDigit := 0

	for j := len(stack) - 1; j > -1; j-- {

		currentChar := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		switch currentChar { //stack[j]
		case "+":
			result += lastDigit
			lastDigit = 0
		case "-":

			//no last digit and result is negative
			if lastDigit == 0 && result < 0 {
				result *= -1
			} else {
				result -= lastDigit
			}

			lastDigit = 0
		case " ":
			continue
		case "(":

			//continue
		default:
			lastDigit, _ = strconv.Atoi(currentChar)
		}
	}

	if lastDigit != 0 {
		result += lastDigit
	}

	return result, stack
}
*/
