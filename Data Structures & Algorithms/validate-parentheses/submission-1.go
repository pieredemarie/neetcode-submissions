func isValid(s string) bool {
    matching := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := make([]rune, 0,len(s))

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)

		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if top != matching[char] {
				return false
			} 
		}
	}

	return len(stack) == 0
}
