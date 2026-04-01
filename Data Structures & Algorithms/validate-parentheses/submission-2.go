func isValid(s string) bool {
    legend := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}

	stack := []string{}
	for i := 0; i < len(s); i++ {
		str := string(s[i])
		switch str {
		case "(", "[", "{":
			stack = append(stack, str)
		case ")", "]", "}":
			if len(stack) == 0 {
				return false
			}

			opener := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop

			if legend[opener] != str {
				return false
			}
		}
	}

	return len(stack) == 0
}
