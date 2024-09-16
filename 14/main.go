package main

func main() {

	strs := []string{"flower", "flow", "flight"}

	//["a","ac"]

	//"flower", "flow", "flight"

	//"flower", "flower", "flower", "flower"

	//["aaa","aa","aaa"]

	// "a"

	// ""

	println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {

	//find smallest word

	length := len(strs[0])

	for j := 0; j < len(strs); j++ {

		if len(strs[j]) < length {
			length = len(strs[j])
		}
	}

	//check each char available in all words, if not stop

	prefix := ""

	for i := 0; i < length; i++ {

		availableInAllWords := true

		for j := 1; j < len(strs); j++ {

			if strs[0][i] != strs[j][i] {
				availableInAllWords = false
				break
			}
		}

		if availableInAllWords {
			prefix += string(strs[0][i])
		} else {
			return prefix
		}
	}

	return prefix
}
