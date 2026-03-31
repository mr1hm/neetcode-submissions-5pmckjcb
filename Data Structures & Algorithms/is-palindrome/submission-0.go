func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}

	str := ""
	for i := 0; i < len(s); i++ {
		if unicode.IsLetter(rune(s[i])) || unicode.IsDigit(rune(s[i])) {
			str += string(s[i])
		}
	}
	str = strings.ToLower(str)
	
	b := []byte(str)
	i, j := 0, len(b)-1
	for i < j {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}

	if string(b) == str {
		return true
	}

	return false
}
