package main

func main() {
	s := "()[]{"

	println(isValid(s))
}

func isValid(s string) bool {
	var a []string

	for i := 0; i < len(s); i++ {

		if isClosingFound(s, i, a) {
			a = a[:len(a)-1]
		} else {
			a = append(a, string(s[i]))
		}
	}

	//if stack is empty, it was valid

	return len(a) == 0
}

/**
 * if closing found in last element of stack
 */
func isClosingFound(s string, i int, a []string) bool {

	pair := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}

	//if opening

	_, exists := pair[string(s[i])]
	if exists {
		return false
	}

	if len(a) == 0 {
		return false
	}

	return pair[a[len(a)-1]] == string(s[i])
}
